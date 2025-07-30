curl -H "Host: ifsapp.kawancn.com" \
    -H "Content-Type: application/json;charset=UTF-8" \
    -H "Origin: http://ifsapp.kawancn.com" \
    -H "Accept: application/json, text/plain, */*" \
    -H "User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148" \
    -H "Referer: http://ifsapp.kawancn.com/Pages/GamePlay/H5Game/game_bg/index.aspx" \
    -H "Accept-Language: zh-CN,zh-Hans;q=0.9" \
    --data-binary "{\"userid\":\"{{ .UserId }}\",\"ptype\":\"1\",\"deviceid\":\"{{ .DeviceId }}\",\"unix\":\"{{ .Unix }}\",\"token\":\"{{ .Token }}\",\"keycode\":\"{{ .KeyCode }}\",\"appversion\":\"123\",\"barheight\":\"48\"}" \
    --compressed "http://ifsapp.kawancn.com/IFS/pc28/pc28_UserBaseH5N.ashx"
