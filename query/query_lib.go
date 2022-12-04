package query

const (
	GetUserAll = "select data.user_id, data.username, data.first_name, data.last_name, data.status, data.created_date, data.password, data.refresh_token, data.user_data\n" +
		"from (select u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, ua1.data_value as password, coalesce(ua2.data_value,'') as refresh_token, concat('{',string_agg(concat('\"',ud.data_key,'\":\"',ud.data_value,'\"'),','),'}') as user_data\n" +
		"from public.users u\n" +
		"left join user_authentication ua1 on u.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'\n" +
		"left join user_authentication ua2 on u.user_id = ua2.user_id and ua2.data_key = 'REFRESH_TOKEN'\n" +
		"left join user_data ud on u.user_id = ud.user_id\n" +
		"group by (u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, password, refresh_token)) as data \n"

	//etUserAll = "select data.user_id, data.username, data.first_name, data.last_name, data.status, data.created_date, data.password, data.refresh_token, data.user_data\n" +
	//	"from (select u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, ua1.data_value as password, coalesce(ua2.data_value,'') as refresh_token,\n" +
	//	"case count(ud.data_key) when 0\n" +
	//	"then ''\n" +
	//	"else concat('{',string_agg(concat('\"',ud.data_key,'\":\"',ud.data_value,'\"'),','),'}')\n" +
	//	"end as user_data\n" +
	//	"from public.users u\n" +
	//	"left join user_authentication ua1 on u.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'\n" +
	//	"left join user_authentication ua2 on u.user_id = ua2.user_id and ua2.data_key = 'REFRESH_TOKEN'\n" +
	//	"left join user_data ud on u.user_id = ud.user_id\n" +
	//	"group by (u.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, password, refresh_token)) as data \n"

	GetUserByEmail = "select ud.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, ua1.data_key as password, ua2.data_key as refresh_token, concat('{',string_agg(concat('\"',ud2.data_key,'\":\"',ud2.data_value,'\"'),','),'}') as user_data\n" +
		"from public.user_data ud\n" +
		"left join users u on u.user_id = ud.user_id\n" +
		"left join user_authentication ua1 on ud.user_id = ua1.user_id and ua1.data_key = 'PASSWORD'\n" +
		"left join user_authentication ua2 on ud.user_id = ua2.user_id and ua2.data_key = 'REFRESH_TOKEN'\n" +
		"left join user_data ud2 on u.user_id = ud.user_id and ud.user_id = ud2.user_id\n" +
		"where ud.data_key = 'EMAIL' and ud.data_value = $1 \n" +
		"group by (ud.user_id, u.username, u.first_name, u.last_name, u.status, u.created_date, password, refresh_token)\n"

	GetListAccount = "select username\nfrom public.users\nunion\nselect data_value\nfrom public.user_data\nwhere data_key = 'EMAIL'"

	GetWalletType = "select wt.wallet_code, wt.wallet_name from public.wallet_type wt;"

	//GetWalletByUserId = "select w.user_id, w.wallet_id, w.wallet_name, w.type from public.wallets w where w.is_active = 'Y' and w.user_id = $1 order by w.wallet_id ASC;"
	GetWalletByUserId = "select w.user_id, w.wallet_id, w.wallet_name, wt.wallet_code from public.wallets w join public.wallet_type wt on wt.wallet_code = w.type where w.is_active = 'Y' and w.user_id = $1 order by w.wallet_id ASC;"

	GetActivityTypes = "select at.activity_type_id, at.type_code, at.type_name, at.category, at.sub_category_name\nfrom public.activity_type at\nwhere at.is_active = 'Y'\norder by at.type_code, at.category ASC"
)
