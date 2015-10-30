// Package sumeragi implements time in Japanese calendar(皇暦, 皇紀, 元号).
package sumeragi

import (
	"time"
)

var SUMERAGI = time.FixedZone("SUMERAGI", 32400)

// Now in Japanese calendar.
func Now() time.Time {
	return time.Now().In(SUMERAGI).AddDate(660, 0, 0)
}

// To time in Japanese calendar.
func To(t time.Time) time.Time {
	if z, _ := t.Zone(); z == "SUMERAGI" {
		return t
	}
	return t.In(SUMERAGI).AddDate(660, 0, 0)
}

var JST = time.FixedZone("JST", 32400)

// Era of Japan in Romaji and Kanji.
func Era(t time.Time) (string, string) {
	if z, _ := t.Zone(); z == "SUMERAGI" {
		t = t.AddDate(-660, 0, 0)
	}
	t = t.In(JST)
	y := t.Year()
	m := t.Month()
	if y < 593 {
		return "", ""
	}
	if y < 629 {
		return "Suiko", "推古"
	}
	if y < 642 {
		return "Jomei", "舒明"
	}
	if y < 645 {
		return "Kougyoku", "皇極"
	}
	if y < 650 {
		return "Taika", "大化"
	}
	if y < 655 {
		return "Hakuchi", "白雉"
	}
	if y < 662 {
		return "Saimei", "斉明"
	}
	if y < 672 {
		return "Tenji", "天智"
	}
	if y < 673 {
		return "Koubun", "弘文"
	}
	if y < 686 {
		return "Tenmu", "天武"
	}
	if y < 687 {
		return "Shuchou", "朱鳥"
	}
	if y < 697 {
		return "Jitou", "持統"
	}
	if y < 701 {
		return "Monmu", "文武"
	}
	if y < 704 {
		return "Taihou", "大宝"
	}
	if y < 708 {
		return "Keiun", "慶雲"
	}
	if y < 715 {
		return "Wadou", "和銅"
	}
	if y < 717 {
		return "Reiki", "霊亀"
	}
	if y < 724 {
		return "Yourou", "養老"
	}
	if y < 729 {
		return "Jinki", "神亀"
	}
	if y < 749 {
		return "Tenpyou", "天平"
	}
	if y < 749 && m < time.April {
		return "TenpyouKanpou", "天平感宝"
	}
	if y < 757 {
		return "TenpyouShouhou", "天平勝宝"
	}
	if y < 765 {
		return "TenpyouHouji", "天平宝字"
	}
	if y < 767 {
		return "TenpyouJingo", "天平神護"
	}
	if y < 770 {
		return "JingoKeiun", "神護景曇"
	}
	if y < 781 {
		return "Houki", "宝亀"
	}
	if y < 782 {
		return "Tenou", "天応"
	}
	if y < 806 {
		return "Enryaku", "延暦"
	}
	if y < 810 {
		return "daidou", "大同"
	}
	if y < 824 {
		return "Kounin", "弘仁"
	}
	if y < 834 {
		return "Tenchou", "天長"
	}
	if y < 848 {
		return "Jouwa", "承和"
	}
	if y < 851 {
		return "Kashou", "嘉祥"
	}
	if y < 854 {
		return "Ninju", "仁寿"
	}
	if y < 857 {
		return "Saikou", "斉衡"
	}
	if y < 859 {
		return "Tenan", "天安"
	}
	if y < 877 {
		return "Jougan", "貞観"
	}
	if y < 885 {
		return "Gangyou", "元慶"
	}
	if y < 889 {
		return "Ninna", "仁和"
	}
	if y < 898 {
		return "Kanpyou", "寛平"
	}
	if y < 901 {
		return "Shoutai", "昌泰"
	}
	if y < 923 {
		return "Engi", "延喜"
	}
	if y < 931 {
		return "Enchou", "延長"
	}
	if y < 938 {
		return "Shouhei", "承平"
	}
	if y < 947 {
		return "Tengyou", "天慶"
	}
	if y < 957 {
		return "Tenryaku", "天暦"
	}
	if y < 961 {
		return "Tentoku", "天徳"
	}
	if y < 964 {
		return "Ouwa", "応和"
	}
	if y < 968 {
		return "Kouhou", "康保"
	}
	if y < 970 {
		return "Annna", "安和"
	}
	if y < 973 {
		return "Tenroku", "天禄"
	}
	if y < 976 {
		return "Tenen", "天延"
	}
	if y < 978 {
		return "Jougen", "貞元"
	}
	if y < 983 {
		return "Tengen", "天元"
	}
	if y < 985 {
		return "Eikan", "永観"
	}
	if y < 987 {
		return "Kanna", "寛和"
	}
	if y < 989 {
		return "Eien", "永延"
	}
	if y < 990 {
		return "Eiso", "永祚"
	}
	if y < 995 {
		return "Shouryaku", "正暦"
	}
	if y < 999 {
		return "Choutoku", "長徳"
	}
	if y < 1004 {
		return "Chouhou", "長保"
	}
	if y < 1012 {
		return "Kankou", "寛弘"
	}
	if y < 1017 {
		return "Chouwa", "長和"
	}
	if y < 1021 {
		return "Kannin", "寛仁"
	}
	if y < 1024 {
		return "Chian", "治安"
	}
	if y < 1028 {
		return "Manju", "万寿"
	}
	if y < 1037 {
		return "Chougen", "長元"
	}
	if y < 1040 {
		return "Chouryaku", "長暦"
	}
	if y < 1044 {
		return "Choukyuu", "長久"
	}
	if y < 1046 {
		return "Kantoku", "寛徳"
	}
	if y < 1053 {
		return "Eishou", "永承"
	}
	if y < 1058 {
		return "Tengi", "天喜"
	}
	if y < 1065 {
		return "Kouhei", "康平"
	}
	if y < 1069 {
		return "Jiryaku", "治暦"
	}
	if y < 1074 {
		return "Enkyuu", "延久"
	}
	if y < 1077 {
		return "Jouhou", "承保"
	}
	if y < 1081 {
		return "Shouryaku", "承暦"
	}
	if y < 1084 {
		return "Eihou", "永保"
	}
	if y < 1087 {
		return "Outoku", "応徳"
	}
	if y < 1094 {
		return "Kanji", "寛治"
	}
	if y < 1096 {
		return "Kahou", "嘉保"
	}
	if y < 1097 {
		return "Eichou", "永長"
	}
	if y < 1099 {
		return "Shoutoku", "承徳"
	}
	if y < 1104 {
		return "Kouwa", "康和"
	}
	if y < 1106 {
		return "Chouji", "長治"
	}
	if y < 1108 {
		return "Kashou", "嘉承"
	}
	if y < 1110 {
		return "Tennin", "天人"
	}
	if y < 1113 {
		return "Tenei", "天永"
	}
	if y < 1118 {
		return "Eikyuu", "永久"
	}
	if y < 1120 {
		return "Genei", "元永"
	}
	if y < 1124 {
		return "Houan", "保安"
	}
	if y < 1126 {
		return "Tenji", "天治"
	}
	if y < 1131 {
		return "Daiji", "大治"
	}
	if y < 1132 {
		return "Tenshou", "天承"
	}
	if y < 1135 {
		return "Choushou", "長承"
	}
	if y < 1141 {
		return "Houen", "保延"
	}
	if y < 1142 {
		return "Eiji", "永治"
	}
	if y < 1144 {
		return "Kouji", "康治"
	}
	if y < 1145 {
		return "Tenyou", "天養"
	}
	if y < 1151 {
		return "Kyuuan", "久安"
	}
	if y < 1154 {
		return "Ninpei", "仁平"
	}
	if y < 1156 {
		return "Kyuuju", "久寿"
	}
	if y < 1159 {
		return "Hougen", "保元"
	}
	if y < 1160 {
		return "heiji", "平治"
	}
	if y < 1161 {
		return "Eiryaku", "永暦"
	}
	if y < 1163 {
		return "Ouhou", "応保"
	}
	if y < 1165 {
		return "Choukan", "長寛"
	}
	if y < 1166 {
		return "Eiman", "永万"
	}
	if y < 1169 {
		return "Ninan", "仁安"
	}
	if y < 1171 {
		return "Kaou", "嘉応"
	}
	if y < 1175 {
		return "Jouan", "承安"
	}
	if y < 1177 {
		return "Angen", "安元"
	}
	if y < 1181 {
		return "Jishou", "治承"
	}
	if y < 1182 {
		return "Youwa", "養和"
	}
	if y < 1184 {
		return "Juei", "寿永"
	}
	if y < 1185 {
		return "Genryaku", "元暦"
	}
	if y < 1190 {
		return "Bunji", "文治"
	}
	if y < 1199 {
		return "Kenkyuu", "建久"
	}
	if y < 1201 {
		return "Shouji", "正治"
	}
	if y < 1204 {
		return "Kennin", "建仁"
	}
	if y < 1206 {
		return "Genkyuu", "元久"
	}
	if y < 1207 {
		return "Kenei", "建永"
	}
	if y < 1211 {
		return "Jougen", "承元"
	}
	if y < 1213 {
		return "Kenryaku", "建暦"
	}
	if y < 1219 {
		return "Kenpou", "建保"
	}
	if y < 1222 {
		return "Joukyuu", "承久"
	}
	if y < 1224 {
		return "Jouou", "貞応"
	}
	if y < 1225 {
		return "Gennin", "元仁"
	}
	if y < 1227 {
		return "Karoku", "嘉禄"
	}
	if y < 1229 {
		return "Antei", "安貞"
	}
	if y < 1232 {
		return "Kanki", "寛喜"
	}
	if y < 1233 {
		return "Jouei", "貞永"
	}
	if y < 1234 {
		return "Tempuku", "天福"
	}
	if y < 1235 {
		return "Bunryaku", "文暦"
	}
	if y < 1238 {
		return "Katei", "嘉禎"
	}
	if y < 1239 {
		return "Ryakunin", "暦仁"
	}
	if y < 1240 {
		return "Enou", "延応"
	}
	if y < 1243 {
		return "Ninji", "仁治"
	}
	if y < 1247 {
		return "Kangen", "寛元"
	}
	if y < 1249 {
		return "Houji", "宝治"
	}
	if y < 1256 {
		return "Kenchou", "建長"
	}
	if y < 1257 {
		return "Kougen", "康元"
	}
	if y < 1259 {
		return "Shouka", "正嘉"
	}
	if y < 1260 {
		return "Shougen", "正元"
	}
	if y < 1261 {
		return "Bunou", "文応"
	}
	if y < 1264 {
		return "Kouchou", "弘長"
	}
	if y < 1275 {
		return "Bunei", "文永"
	}
	if y < 1278 {
		return "Kenji", "健治"
	}
	if y < 1288 {
		return "Kouan", "弘安"
	}
	if y < 1293 {
		return "Shouou", "正応"
	}
	if y < 1299 {
		return "Einin", "永仁"
	}
	if y < 1302 {
		return "Shouan", "正安"
	}
	if y < 1303 {
		return "Kangen", "乾元"
	}
	if y < 1306 {
		return "Kangen", "嘉元"
	}
	if y < 1308 {
		return "Tokuji", "徳治"
	}
	if y < 1311 {
		return "Enkei", "延慶"
	}
	if y < 1312 {
		return "Ouchou", "応長"
	}
	if y < 1317 {
		return "Shouwa", "正和"
	}
	if y < 1319 {
		return "Bunpou", "文保"
	}
	if y < 1321 {
		return "Genou", "元応"
	}
	if y < 1324 {
		return "Genkou", "元亨"
	}
	if y < 1426 {
		return "Shouchuu", "正中"
	}
	if y < 1329 {
		return "Karyaku", "嘉暦"
	}
	if y < 1331 {
		return "Gentoku", "元徳"
	}
	if y < 1332 {
		return "Genkou", "元弘"
	}
	if y < 1334 {
		return "Shoukei", "正慶"
	}
	if y < 1336 {
		return "Kenmu", "建武"
	}
	if y < 1338 {
		return "Engen", "延元"
	}
	if y < 1340 {
		return "Ryakuou", "暦応"
	}
	if y < 1342 {
		return "Koukoku", "興国"
	}
	if y < 1345 {
		return "Kouei", "康永"
	}
	if y < 1346 {
		return "Jouwa", "貞和"
	}
	if y < 1350 {
		return "Shouhei", "正平"
	}
	if y < 1352 {
		return "Kanou", "観応"
	}
	if y < 1356 {
		return "Bunwa", "文和"
	}
	if y < 1361 {
		return "Enbun", "延文"
	}
	if y < 1362 {
		return "kouan", "康安"
	}
	if y < 1368 {
		return "Jouji", "貞治"
	}
	if y < 1370 {
		return "Ouan", "応安"
	}
	if y < 1372 {
		return "Kentoku", "建徳"
	}
	if y < 1375 {
		return "Bunchuu", "文中"
	}
	if y < 1381 {
		return "Tenju", "天授"
	}
	if y < 1384 {
		return "Kouwa", "弘和"
	}
	if y < 1390 {
		return "Genchuu", "元中"
	}
	if y < 1394 {
		return "Meitoku", "明徳"
	}
	if y < 1428 {
		return "Ouei", "応永"
	}
	if y < 1429 {
		return "Shouchou", "正長"
	}
	if y < 1441 {
		return "Eikyou", "永享"
	}
	if y < 1444 {
		return "Kakitsu", "嘉吉"
	}
	if y < 1449 {
		return "Bunnan", "文安"
	}
	if y < 1452 {
		return "Houtoku", "宝徳"
	}
	if y < 1455 {
		return "Kyoutoku", "享徳"
	}
	if y < 1457 {
		return "Koushou", "康正"
	}
	if y < 1460 {
		return "Chouroku", "長禄"
	}
	if y < 1466 {
		return "Kanshou", "寛正"
	}
	if y < 1467 {
		return "Bunshou", "文正"
	}
	if y < 1469 {
		return "Ounin", "応仁"
	}
	if y < 1487 {
		return "Bunmei", "文明"
	}
	if y < 1489 {
		return "Choukyou", "長享"
	}
	if y < 1492 {
		return "Entoku", "延徳"
	}
	if y < 1501 {
		return "Meiou", "明応"
	}
	if y < 1504 {
		return "Bunki", "文亀"
	}
	if y < 1521 {
		return "Eishou", "永正"
	}
	if y < 1528 {
		return "Taiei", "大永"
	}
	if y < 1532 {
		return "Kyouroku", "享禄"
	}
	if y < 1555 {
		return "Tenbun", "天文"
	}
	if y < 1558 {
		return "Kouji", "弘治"
	}
	if y < 1570 {
		return "Eiroku", "永禄"
	}
	if y < 1573 {
		return "Genki", "元亀"
	}
	if y < 1592 {
		return "Tenshou", "天正"
	}
	if y < 1596 {
		return "Bunroku", "文禄"
	}
	if y < 1615 {
		return "Keichou", "慶長"
	}
	if y < 1624 {
		return "Genna", "元和"
	}
	if y < 1644 {
		return "Kanei", "寛永"
	}
	if y < 1648 {
		return "Shouhou", "正保"
	}
	if y < 1652 {
		return "Keian", "慶安"
	}
	if y < 1655 {
		return "Shouou", "承応"
	}
	if y < 1658 {
		return "Meireki", "明暦"
	}
	if y < 1661 {
		return "Manji", "万治"
	}
	if y < 1673 {
		return "Kanbun", "寛文"
	}
	if y < 1681 {
		return "Enpou", "延宝"
	}
	if y < 1684 {
		return "Tenwa", "天和"
	}
	if y < 1688 {
		return "Joukyou", "貞享"
	}
	if y < 1704 {
		return "Genroku", "元禄"
	}
	if y < 1711 {
		return "Houei", "宝永"
	}
	if y < 1716 {
		return "Shoutoku", "正徳"
	}
	if y < 1736 {
		return "Kyouhou", "享保"
	}
	if y < 1741 {
		return "Genbun", "元文"
	}
	if y < 1744 {
		return "Kanpou", "寛保"
	}
	if y < 1748 {
		return "Enkyou", "延享"
	}
	if y < 1751 {
		return "Kanen", "寛延"
	}
	if y < 1764 {
		return "Houreki", "宝暦"
	}
	if y < 1772 {
		return "Meiwa", "明和"
	}
	if y < 1781 {
		return "Anei", "安永"
	}
	if y < 1789 {
		return "Tenmei", "天明"
	}
	if y < 1801 {
		return "Kansei", "寛政"
	}
	if y < 1804 {
		return "Kyouwa", "享和"
	}
	if y < 1818 {
		return "Bunka", "文化"
	}
	if y < 1830 {
		return "Bunsei", "文政"
	}
	if y < 1844 {
		return "Tempou", "天保"
	}
	if y < 1848 {
		return "Kouka", "弘化"
	}
	if y < 1854 {
		return "Kaei", "嘉永"
	}
	if y < 1860 {
		return "Ansei", "安政"
	}
	if y < 1861 {
		return "Manen", "万延"
	}
	if y < 1864 {
		return "Bunkyuu", "文久"
	}
	if y < 1865 {
		return "Genji", "元治"
	}
	if y < 1868 {
		return "Keiou", "慶応"
	}
	if y < 1912 {
		return "Meiji", "明治"
	}
	if y < 1926 {
		return "Taishou", "大正"
	}
	if y < 1989 {
		return "Shouwa", "昭和"
	}
	return "Heisei", "平成"
}

// TODO(pullreq welcome): 祭日
