/**
 * Created by igor on 10.07.17.
 */

"use strict";

const log = require(__appRoot + '/lib/log')(module),
    CodeError = require(__appRoot + '/lib/error'),
    buildQuery = require('./utils').buildQuery;

const create = `
    INSERT INTO callback_queue (name, domain, description, agents)
    VALUES ($1, $2, $3, $4)
    RETURNING *;             
`;

const sqlItem = `
    SELECT * FROM callback_queue WHERE id = $1 AND domain = $2
`;

const sqlDelete = `
    DELETE FROM callback_queue WHERE id = $1 AND domain = $2
    RETURNING id;
`;

const sqlMemberCreate = `
    INSERT INTO callback_members (domain, queue_id, number, href, user_agent, location)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *, (SELECT name FROM callback_queue WHERE id = $2) as queue_name;  
`;


const sqlMemberItem = `
    SELECT 
        m.*, 
        w.name as widget_name, 
        (SELECT array_to_json(array_agg(c)) FROM callback_members_comment as c where c.member_id = m.id) as comments,
        (SELECT array_to_json(array_agg(cl)) FROM callback_calls as cl where cl.member_id = m.id LIMIT 20) as calls
    FROM callback_members as m
     LEFT JOIN widget as w on w.id = m.widget_id
    WHERE m.id = $1 AND m.queue_id = $2 AND m.domain = $3
`;

const sqlMemberDelete = `
    DELETE FROM callback_members WHERE id = $1 AND queue_id = $2 AND domain = $3
    RETURNING *;
`;

const sqlCommentAdd = `
    INSERT INTO callback_members_comment (created_by, created_on, member_id, text)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_by, created_on::bigint, member_id, text, (
        SELECT row_to_json(d)
        FROM (
          SELECT
            q.name AS queue_name, m.widget_id AS widget_id, q.id AS queue_id
          FROM callback_members m
            JOIN callback_queue q ON q.id = m.queue_id
          WHERE m.id = $3
          LIMIT 1
        ) d
    ) as info;
`;
const sqlCommentRemove = `
    DELETE FROM callback_members_comment WHERE id = $1 AND member_id = $2
    RETURNING *;
`;
const sqlCommentUpdate = `
    UPDATE callback_members_comment SET text = $3  WHERE id = $1 AND member_id = $2
    RETURNING *;
`;

function add(pool) {
    return {
        list: (request, cb) => {
            buildQuery(pool, request, "callback_queue", cb);
        },

        findById: (_id, domainName, cb) => {
            pool.query(
                sqlItem,
                [
                    +_id,
                    domainName
                ], (err, res) => {
                    if (err) {
                        return cb(err);
                    }
                    if (res && res.rowCount) {
                        return cb(null, res.rows[0])
                    } else {
                        return cb(new CodeError(404, `Not found ${_id}@${domainName}`));
                    }
                }
            )
        },

        create: (doc, cb) => {
            pool.query(
                create,
                [
                    doc.name, //$1
                    doc.domain, //$2
                    doc.description, //$3
                    doc.agents ? doc.agents : [], //$3
                ], (err, res) => {
                    if (err)
                        return cb(err);

                    if (res && res.rowCount) {
                        return cb(null, res.rows[0])
                    } else {
                        log.error('bad response', res);
                        return cb(new Error('Bad db response'));
                    }
                }
            )
        },

        delete: (id, domain, cb) => {
            pool.query(
                sqlDelete,
                [
                    id,
                    domain
                ], (err, res) => {
                    if (err)
                        return cb(err);

                    if (res && res.rowCount) {
                        return cb(null, res.rows[0].id, res.rows[0]._filepath)
                    } else {
                        return cb(new CodeError(404, `Not found ${id}@${domain}`));
                    }
                }
            )
        },

        update: (_id, domainName, doc = {}, cb) => {
            const values = [];
            const params = [];


            for (let field of allowQueueUpdateFields) {
                if (doc.hasOwnProperty(field)) {
                    values.push(`${field} = $` + params.push(doc[field]));
                }
            }

            let update = `
                    UPDATE callback_queue 
                        SET ${values.join(',')} 
                    WHERE id = $${params.length + 1} AND domain = $${params.length + 2}
                    RETURNING *
                `;
            params.push(+_id);
            params.push(domainName);
            pool.query(
                update,
                params, (err, res) => {
                    if (err)
                        return cb(err);

                    if (res && res.rowCount) {
                        return cb(null, res.rows)
                    } else {
                        return cb(new CodeError(404, `Not found ${_id}@${domainName}`));
                    }
                }
            );
        },

        members: {
            addComment: (member_id, data, cb) => {
                pool.query(
                    sqlCommentAdd, //created_by, created_on, member_id, text
                    [
                        data.created_by, //$1
                        +data.created_on, //$2
                        +member_id, //$3
                        data.text, //$4
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            log.error('bad response', res);
                            return cb(new Error('Bad db response'));
                        }
                    }
                )
            },
            removeComment: (id, queue_id, domain, comment_id, cb) => {
                pool.query(
                    sqlCommentRemove,
                    [
                        +comment_id, //$1
                        +id, //$2
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            log.error('bad response', res);
                            return cb(new Error('Bad db response'));
                        }
                    }
                )
            },
            updateComment: (id, queue_id, domain, comment_id, text, cb) => {
                pool.query(
                    sqlCommentUpdate,
                    [
                        +comment_id, //$1
                        +id, //$2
                        text
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            log.error('bad response', res);
                            return cb(new Error('Bad db response'));
                        }
                    }
                )
            },
            list: (request, cb) => {
                buildQuery(pool, request, "callback_members", cb);
            },

            viewIsOverdue: (userId, domain, limit = 40, offset = 0, filter, cb) => {
                const params = [
                    domain,
                    Date.now(),
                    limit,
                    offset,
                    userId
                ];
                if (filter) {
                    params.push(filter)
                }
                pool.query(
                    `select
                      m.id,
                      m.number,
                      m.callback_time,
                      m.done,
                      c2.name as queue_name,
                      c2.id as queue_id,
                      (extract(EPOCH from m.created_on) * 1000)::int8 created_on
                    from callback_members m
                      inner join callback_queue c2 on m.queue_id = c2.id and c2.agents @> ARRAY[$5]::VARCHAR(100)[]
                    where m.domain = $1 and not done is true and callback_time < $2 ${filter ? "AND number like $6 || '%'" : ""}
                    order by callback_time desc
                    limit $3 offset $4`,
                    params,
                    (err, res) => {
                        if (err) {
                            return cb(err);
                        }
                        return cb(null, res.rows)
                    }
                )
            },

            viewIsScheduled: (userId, domain, limit = 40, offset = 0, filter, cb) => {
                const params = [
                    domain,
                    Date.now(),
                    limit,
                    offset,
                    userId
                ];

                if (filter) {
                    params.push(filter)
                }

                pool.query(
                    `select
                      m.id,
                      m.number,
                      m.callback_time,
                      m.done,
                      c2.name as queue_name,
                      c2.id as queue_id,
                      (extract(EPOCH from m.created_on) * 1000)::int8 created_on
                    from callback_members m
                      inner join callback_queue c2 on m.queue_id = c2.id and c2.agents @> ARRAY[$5]::VARCHAR(100)[]
                    where m.domain = $1 and not done is true and callback_time >= $2 ${filter ? "AND number like $6 || '%'" : ""}
                    order by callback_time asc
                    limit $3 offset $4`,
                    params,
                    (err, res) => {
                        if (err) {
                            return cb(err);
                        }
                        return cb(null, res.rows)
                    }
                )
            },

            viewIsCompleted: (userId, domain, limit = 40, offset = 0, filter, cb) => {
                const params = [
                    domain,
                    limit,
                    offset,
                    userId
                ];

                if (filter) {
                    params.push(filter)
                }

                pool.query(
                    `select
                      m.id,
                      m.number,
                      m.callback_time,
                      m.done,
                      m.done_at,
                      c2.name as queue_name,
                      c2.id as queue_id,
                      (extract(EPOCH from m.created_on) * 1000)::int8 created_on
                    from callback_members m
                      inner join callback_queue c2 on m.queue_id = c2.id and c2.agents @> ARRAY[$4]::VARCHAR(100)[]
                    where m.domain = $1 and done is true ${filter ? "AND number like $5 || '%'" : ""}
                    order by done_at desc
                    limit $2 offset $3`,
                    params,
                    (err, res) => {
                        if (err) {
                            return cb(err);
                        }
                        return cb(null, res.rows)
                    }
                )
            },

            viewIsNotCallbackTime: (userId, domain, limit = 40, offset = 0, filter, cb) => {
                const params = [
                    domain,
                    limit,
                    offset,
                    userId
                ];

                if (filter) {
                    params.push(filter)
                }

                pool.query(
                    `select
                      m.id,
                      m.number,
                      m.callback_time,
                      m.done,
                      c2.name as queue_name,
                      c2.id as queue_id,
                      (extract(EPOCH from m.created_on) * 1000)::int8 created_on
                    from callback_members m
                      inner join callback_queue c2 on m.queue_id = c2.id and c2.agents @> ARRAY[$4]::VARCHAR(100)[]
                    where m.domain = $1 and not done is true and callback_time is null ${filter ? "AND number like $5 || '%'" : ""}
                    order by created_on asc
                    limit $2 offset $3`,
                    params,
                    (err, res) => {
                        if (err) {
                            return cb(err);
                        }
                        return cb(null, res.rows)
                    }
                )
            },

            findById: (_id, queue_id, domainName, cb) => {
                pool.query(
                    sqlMemberItem,
                    [
                        +_id,
                        +queue_id,
                        domainName
                    ], (err, res) => {
                        if (err) {
                            return cb(err);
                        }
                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            return cb(new CodeError(404, `Not found ${_id}@${domainName}`));
                        }
                    }
                )
            },

            create: (doc, cb) => {
                pool.query(
                    sqlMemberCreate, //domain, queue_id, number, href, user_agent, location
                    [
                        doc.domain, //$1
                        doc.queue_id, //$2
                        doc.number, //$3
                        doc.href, //$4
                        doc.user_agent, //$5
                        doc.location, //$6
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            log.error('bad response', res);
                            return cb(new Error('Bad db response'));
                        }
                    }
                )
            },

            createPublic: (widget_id, doc, cb) => {
                pool.query(
                    `SELECT * FROM insert_member_public($1, $2)`,
                    [
                        +widget_id,
                        doc
                    ],
                    (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            if (res.rows[0].error) {
                                let errText = null;
                                switch (res.rows[0].error) {
                                    case -1:
                                        errText = `Yo IP (${doc.request_ip}) banned!`;
                                        break;
                                    case -2:
                                        errText = `Max call per IP (${doc.request_ip}).`;
                                        break;
                                    case -3:
                                        errText = `Max call per number (${doc.number}).`;
                                        break;
                                    default:
                                        errText = "Unknown"

                                }
                                return cb(new Error(errText))
                            }
                            if (!res.rows[0].destination_number)
                                return cb(new Error('No found destination number'));

                            return cb(null, {
                                member: res.rows[0].member,
                                destinationNumber: res.rows[0].destination_number,
                                queueName: res.rows[0].queue_name,
                                callTimeout: res.rows[0].call_timeout
                            })
                        }
                        return cb(new Error('Bad db response'));
                    }
                )
            },

            delete: (id, queue_id, domain, cb) => {
                pool.query(
                    sqlMemberDelete,
                    [
                        +id,
                        +queue_id,
                        domain
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            return cb(new CodeError(404, `Not found ${id}@${domain}`));
                        }
                    }
                )
            },

            update: (_id, queue_id, domainName, doc = {}, cb) => {
                const values = [];
                const params = [];


                for (let field of allowMemberUpdateFields) {
                    if (doc.hasOwnProperty(field)) {
                        values.push(`${field} = $` + params.push(doc[field]));
                    }
                }

                let update = `
                    UPDATE callback_members m
                        SET ${values.join(',')} 
                    FROM callback_queue q
                    WHERE (m.done is null or m.done = false) AND m.id = $${params.length + 1} AND m.queue_id = $${params.length + 2} AND m.domain = $${params.length + 3} AND m.queue_id = q.id
                    RETURNING m.*, q.name as queue_name
                `;
                params.push(+_id);
                params.push(+queue_id);
                params.push(domainName);
                pool.query(
                    update,
                    params, (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            return cb(new CodeError(404, `Not found ${_id}@${domainName}`));
                        }
                    }
                );
            },

            createCall: (member_id, user, cb) => {
                pool.query(
                    `with ins as (
                    insert into callback_calls (member_id, created_at, created_by)
                    select m.id, $2, $3
                    from callback_members m
                      inner join callback_queue c2 on m.queue_id = c2.id
                      where m.id = $1 and c2.agents && '{${user}}'
                    returning id
                    )
                    select m.number, c2.name as queue_name from callback_members m
                    inner join callback_queue c2 on m.queue_id = c2.id
                      where m.id = $1 and exists(select 1 from ins)`, //domain, queue_id, number, href, user_agent, location
                    [
                        +member_id,
                        Date.now(),
                        user
                    ], (err, res) => {
                        if (err)
                            return cb(err);

                        if (res && res.rowCount) {
                            return cb(null, res.rows[0])
                        } else {
                            log.error('bad response', res);
                            return cb(new Error('Bad db response'));
                        }
                    }
                )
            }
        }
    }
}

module.exports = add;

const allowMemberUpdateFields = ['number', 'href', 'user_agent', 'location', 'domain', 'done', 'done_at', 'done_by'];
const allowQueueUpdateFields = ['name', 'description', 'agents'];
