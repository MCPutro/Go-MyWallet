package query

const (
	GetUserAll = `select data.user_id, data.username, data.first_name, data.last_name, data.status, data.created_date, data.password, data.refresh_token, data.user_data
		from (select u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, ua1.data_value as password, coalesce(ua2.data_value,'') as refresh_token, concat('{',string_agg(concat('"',ud.data_key,'":"',ud.data_value,'"'),','),'}') as user_data
		from public.users u
		left join user_authentication ua1 on u.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'
		left join user_authentication ua2 on u.user_id = ua2.user_id and ua2.data_key = 'REFRESH_TOKEN'
		left join user_data ud on u.user_id = ud.user_id
		group by (u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, password, refresh_token)) as data `

	GetUserByEmail = `select ud.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, ua1.data_value as password, coalesce(ua2.data_value,'') as refresh_token, concat('{',string_agg(concat('"',ud2.data_key,'":"',ud2.data_value,'"'),','),'}') as user_data
		from public.user_data ud
		left join users u on u.user_id = ud.user_id
		left join user_authentication ua1 on ud.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'
		left join user_authentication ua2 on ud.user_id = ua2.user_id and ua2.data_key = 'REFRESH_TOKEN'
		left join user_data ud2 on u.user_id = ud.user_id and ud.user_id = ud2.user_id
		where ud.data_key = 'EMAIL' and ud.data_value = $1 
		group by (ud.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, password, refresh_token)`

	GetListAccount = "select username\nfrom public.users\nunion\nselect data_value\nfrom public.user_data\nwhere data_key = 'EMAIL'"

	GetWalletType = "select wt.wallet_code, wt.wallet_name from public.wallet_type wt;"

	GetWalletById = "select w.user_id, w.wallet_id, w.wallet_name, wt.wallet_name, w.amount from public.wallets w join public.wallet_type wt on wt.wallet_code = w.type " +
		"where w.is_active = 'Y' and %s order by w.wallet_id ASC ;" //w.user_id = $1

	GetActivityTypes = "select ac.category_id, ac.type, ac.category, ac.sub_category_name from public.activity_category ac %s"

	//GetActivityList = "SELECT t.activity_id, t.user_id, t.wallet_id_from, t.wallet_id_to, t.period, t.activity_date, a.activity_type_name, ac.sub_category_name " +
	//	"FROM public.user_activity t " +
	//	"inner join activity_category ac on ac.category_id = t.category_id " +
	//	"inner join activity_type a on a.activity_type_code = ac.type" +
	//	"where %s " +
	//	"ORDER BY t.activity_date DESC ;"
)
