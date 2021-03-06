package wx

const (
	//=======【基本信息设置】=====================================
	//微信公众号身份的唯一标识。审核通过后，在微信发送的邮件中查看
	//或者微信小程序的appid，可在小程序后台设置－开发设置内查看
	appid = ""
	//受理商id，身份标识
	mchid = ""
	//商户支付密钥key。审核通过后，在微信发送的邮件中查看
	key = ""
	//jsapi接口中获取openid，审核后在公众平台开启开发模式后可查看
	//或者微信小程序的appsecret，可在小程序后台设置－开发设置内查看
	appsecret = ""

	//=======【jsapi路径设置】===================================
	//获取access_token过程中的跳转uri，通过跳转将code传入jsapi支付页面
	jsapi_redirect_url = "http://www.example.com/wxpay/demo?type=WX_JSAPI"

	//=======【curl超时设置】===================================
	//本例程通过curl使用http post方法，此处可修改其超时时间，默认为30秒
	curl_timeout = 30
)
