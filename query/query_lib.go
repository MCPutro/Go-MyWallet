package query

const (
	GetUserAll = `select data.user_id,  data.account_id, data.username, data.full_name, data.status, data.created_date, data.password, data.user_data
		from (select u.user_id, u.account_id, u.username, u.full_name, u.status, u.created_date, ua1.data_value as password,  concat('{',string_agg(concat('"',ud.data_key,'":"',ud.data_value,'"'),','),'}') as user_data
      			from public.users u
               	inner join user_authentication ua1 on u.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'
               	inner join user_data ud on u.user_id = ud.user_id
      		group by (u.user_id, ua1.data_value)) as data `

	GetUserByEmail = `select ud.user_id, u.account_id, u.username, u.full_name, u.status, u.created_date, ua1.data_value as password, concat('{',string_agg(concat('"',ud2.data_key,'":"',ud2.data_value,'"'),','),'}') as user_data
		from public.user_data ud
		inner join users u on u.user_id = ud.user_id
		inner join user_authentication ua1 on ud.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'
		inner join user_data ud2 on u.user_id = ud.user_id and ud.user_id = ud2.user_id
		where ud.data_key = 'EMAIL' and ud.data_value = $1 
		group by (ud.user_id, u.account_id, u.username, u.full_name, u.status, u.created_date, password) `

	GetListAccount = `select username from public.users 
                union 
                select data_value from public.user_data where data_key = 'EMAIL' `

	GetWalletType = "select wt.wallet_code, wt.wallet_name from public.wallet_type wt;"

	GetWalletById = `select w.user_id, w.wallet_id, w.wallet_name, wt.wallet_name, w.amount 
		from public.wallets w join public.wallet_type wt on wt.wallet_code = w.type  
		where w.is_active = 'Y' and %s order by w.wallet_id ASC ;`  //w.user_id = $1

	GetActivityTypes = `select ac.category_id, ac.type, ac.category, ac.sub_category_name from public.activity_category ac %s`
)
