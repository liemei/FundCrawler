package models

// 基金基本信息
type FundInfo struct {
	Id int64 `xorm:"autoincr pk 'Id'"`
	Code string `xorm:"varchar(64) 'code' comment('基金代码')"` //基金代码
	Name string `xorm:"varchar(256) 'name' comment('基金名称')"` //基金名称
	FundType string `xorm:"varchar(126) 'fundtype' comment('基金类型')"` //基金类型
	GradeThree string `xorm:"varchar(256) 'gradethree' comment('基金评级3年')"` //基金评级(3年)
	GradeFive string `xorm:"varchar(256) 'gradefive' comment('基金评级5年')"` //基金评级(5年)
	ValueTime string `xorm:"varchar(64)"` //基金价值日期
	ValueUnit string `xorm:"varchar(64)"` //基金价值
	RateOfReturn string `xorm:"varchar(64)"` //今年回报率
	EstablishTime string `xorm:"varchar(64)"` //创建时间
	TotalAssets float64 //总净资产
	FundCompany string `xorm:"varchar(256) 'name' comment('基金公司')"`
	Investment float32 //最低投资额
	Url string `xorm:"varchar(256) 'url' comment('基金链接')"` //基金链接
}
// 历史业绩
type Achievement struct {
	Id int64 `xorm:"autoincr pk 'Id'"`
	Code string `xorm:"varchar(64) 'code' comment('基金代码')"` //基金代码
	Name string `xorm:"varchar(256) 'name' comment('基金名称')"` //基金名称
	TotalRateOfReturn float64 //总回报
	Benchmark float64 //基准指数
	Tlverage float64 //同类平均
	Year int //年份
}
//费率
type Fundrate struct {
	Id int64 `xorm:"autoincr pk 'Id'"`
	Code string `xorm:"varchar(64) 'code' comment('基金代码')"` //基金代码
	Name string `xorm:"varchar(256) 'name' comment('基金名称')"` //基金名称
	ManageRate  string `xorm:"varchar(64) 'ManageRate' comment('管理费率')"` //管理费率
	Trusteeship string `xorm:"varchar(64) 'Trusteeship' comment('托管费')"` //托管费
	ServiceRate string `xorm:"varchar(64) 'ServiceRate' comment('销售服务费')"` //销售服务费
	Subscription string `xorm:"varchar(256) 'Subscription' comment('申购费')"` //申购费
	Redemption string `xorm:"varchar(256) 'Redemption' comment('赎回费')"` //赎回费
}

// 基金风险
type FundRisk struct {
	Id int64 `xorm:"autoincr pk 'Id'"`
	Code string `xorm:"varchar(64) 'code' comment('基金代码')"` //基金代码
	Name string `xorm:"varchar(256) 'name' comment('基金名称')"` //基金名称
	Sharpe float32 //夏普比率
	Standard float32 //标准差
	AverageReturn float32 //平均回报
	Alpha float32 //阿尔法系数
	Alpha1 float32 //阿尔法系数同类平均
	Beta float32 //贝塔系数
	Beta1 float32 //贝塔系数同类平均
	Rsquare float32 //R平方值
	Rsquare1 float32 //R平方值同类平均
}

// 基金经理
type FundManage struct {
	Id int64 `xorm:"autoincr pk 'Id'"`
	Code string `xorm:"varchar(64) 'code' comment('基金代码')"` //基金代码
	Name string `xorm:"varchar(256) 'name' comment('基金名称')"` //基金名称
	ManageName string `xorm:"varchar(64) 'ManageName' comment('基金经理名称')"` //基金经理名称
	Manageintroduce string  `xorm:"varchar(1024) 'Manageintroduce' comment('简介')"` //简介
	ManageTime string `xorm:"varchar(64) 'ManageTime' comment('管理时间')"` //管理时间
	Times string `xorm:"varchar(126) 'Times' comment('时间范围')"` //时间范围
	Score float32 //评分
	Star float32 //星级
}