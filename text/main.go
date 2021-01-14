package text

import plug "github.com/F88888/opsoa_plug"

func main() {
	var plugInfo = plug.Plug{}
	if err := plugInfo.Start("start"); err == nil {
		plugInfo.Set("第N步", "执行成功")
	}
}
