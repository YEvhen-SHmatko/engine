package sqlstore

import (
	"fmt"
	"github.com/webitel/engine/model"
	"github.com/webitel/engine/store"
)

type SqlBucketInQueueStore struct {
	SqlStore
}

func NewSqlBucketInQueueStore(sqlStore SqlStore) store.BucketInQueueStore {
	us := &SqlBucketInQueueStore{sqlStore}
	return us
}

func (s SqlBucketInQueueStore) Create(queueBucket *model.QueueBucket) (*model.QueueBucket, *model.AppError) {
	var out *model.QueueBucket
	if err := s.GetMaster().SelectOne(&out, `with q as (
		insert into cc_bucket_in_queue (queue_id, ratio, bucket_id)
		values (:QueueId, :Ratio, :BucketId)
		returning *
	)
	select q.id, q.ratio, cc_get_lookup(cb.id, cb.name::text) as bucket
	from q
		inner join cc_bucket cb on q.bucket_id = cb.id`,
		map[string]interface{}{
			"QueueId":  queueBucket.QueueId,
			"Ratio":    queueBucket.Ratio,
			"BucketId": queueBucket.Bucket.Id,
		}); nil != err {
		return nil, model.NewAppError("SqlBucketInQueueStore.Save", "store.sql_queue_bucket.save.app_error", nil,
			fmt.Sprintf("queue_id=%v bucket_id=%v, %v", queueBucket.QueueId, queueBucket.Bucket.Id, err.Error()), extractCodeFromErr(err))
	} else {
		return out, nil
	}
}

func (s SqlBucketInQueueStore) Get(domainId, queueId, id int64) (*model.QueueBucket, *model.AppError) {
	var queueBucket *model.QueueBucket
	if err := s.GetReplica().SelectOne(&queueBucket, `select q.id, q.queue_id, q.ratio, cc_get_lookup(cb.id, cb.name::text) as bucket
		from cc_bucket_in_queue q
			inner join cc_bucket cb on q.bucket_id = cb.id
		where q.id = :Id and q.queue_id = :QueueId and cb.domain_id = :DomainId`, map[string]interface{}{
		"Id":       id,
		"DomainId": domainId,
		"QueueId":  queueId,
	}); err != nil {
		return nil, model.NewAppError("SqlBucketInQueueStore.Get", "store.sql_queue_bucket.get.app_error", nil,
			fmt.Sprintf("Id=%v, %s", id, err.Error()), extractCodeFromErr(err))
	} else {
		return queueBucket, nil
	}
}

func (s SqlBucketInQueueStore) GetAllPage(domainId, queueId int64, search *model.SearchQueueBucket) ([]*model.QueueBucket, *model.AppError) {
	var out []*model.QueueBucket

	if _, err := s.GetReplica().Select(&out,
		`select q.id, q.ratio, cc_get_lookup(cb.id, cb.name::text) as bucket
				from cc_bucket_in_queue q
					inner join cc_bucket cb on q.bucket_id = cb.id
				where q.queue_id = :QueueId and cb.domain_id = :DomainId
					and ( (:Q::varchar isnull or (cb.name ilike :Q::varchar ) ))
			order by q.id
			limit :Limit
			offset :Offset`, map[string]interface{}{
			"DomainId": domainId,
			"Limit":    search.GetLimit(),
			"Offset":   search.GetOffset(),
			"Q":        search.GetQ(),
			"QueueId":  queueId,
		}); err != nil {
		return nil, model.NewAppError("SqlBucketInQueueStore.GetAllPage", "store.sql_queue_bucket.get_all.app_error",
			nil, err.Error(), extractCodeFromErr(err))
	} else {
		return out, nil
	}
}

func (s SqlBucketInQueueStore) Update(domainId int64, queueBucket *model.QueueBucket) (*model.QueueBucket, *model.AppError) {
	err := s.GetMaster().SelectOne(&queueBucket, `with q as (
		update cc_bucket_in_queue bq
			set ratio = :Ratio,
				bucket_id = :BucketId
		from cc_queue cq
		where bq.id = :Id and cq.id = :QueueId and cq.domain_id = :DomainId
		returning bq.*
	)
	select q.id, q.ratio, cc_get_lookup(cb.id, cb.name::text) as bucket
	from q
		inner join cc_bucket cb on q.bucket_id = cb.id`, map[string]interface{}{
		"Ratio":    queueBucket.Ratio,
		"BucketId": queueBucket.Bucket.Id,
		"Id":       queueBucket.Id,
		"QueueId":  queueBucket.QueueId,
		"DomainId": domainId,
	})
	if err != nil {
		return nil, model.NewAppError("SqlBucketInQueueStore.Update", "store.sql_queue_bucket.update.app_error", nil,
			fmt.Sprintf("Id=%v, %s", queueBucket.Id, err.Error()), extractCodeFromErr(err))
	}
	return queueBucket, nil
}

func (s SqlBucketInQueueStore) Delete(queueId, id int64) *model.AppError {
	if _, err := s.GetMaster().Exec(`delete from cc_bucket_in_queue c where c.id=:Id and c.queue_id = :QueueId`,
		map[string]interface{}{"Id": id, "QueueId": queueId}); err != nil {
		return model.NewAppError("SqlBucketInQueueStore.Delete", "store.sql_queue_bucket.delete.app_error", nil,
			fmt.Sprintf("Id=%v, %s", id, err.Error()), extractCodeFromErr(err))
	}
	return nil
}
