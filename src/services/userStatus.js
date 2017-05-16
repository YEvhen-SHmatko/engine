/**
 * Created by Igor Navrotskyj on 16.09.2015.
 */

'use strict';

var CodeError = require(__appRoot + '/lib/error'),
    log = require(__appRoot + '/lib/log')(module),
    conf = require(__appRoot + '/conf')
    ;

const noWriteStatus = String(conf.get('application:writeUserStatus')) !== 'true';

var Service = {
    insert: function (option) {
        if (noWriteStatus) return;

        if (!option['state'] || !option['status'] || !option['account']) {
            return log.warn('Caller %s status or state undefined.', option['account']);
        }

        let data = option;
        data['date'] = Date.now();
        let dbUserStatus = application.DB._query.userStatus;

        dbUserStatus.create(data, (err) => {
            if (err)
                log.error(err);
        });
    },
    
    _removeByUserId: function (userId, domain, cb) {
        if (!domain || !userId) {
            return cb(new CodeError(400, "Domain is required."));
        };

        var dbUserStatus = application.DB._query.userStatus;
        return dbUserStatus._removeByUserId(domain, userId, cb);
    }
};

module.exports = Service;