package pv

import "net/url"

// Request Header
var (
	referer    = "https://www.pixiv.net/"
	cookie     = "first_visit_datetime_pc=2022-04-23+19%3A55%3A32; p_ab_id=7; p_ab_id_2=7; p_ab_d_id=1615310070; yuid_b=JEeBGEI; _fbp=fb.1.1650711335517.669551962; _ga=GA1.2.1541353225.1650711335; device_token=ba1b528d6519178761605d39f698183f; privacy_policy_agreement=3; c_type=20; privacy_policy_notification=0; a_type=0; b_type=1; _im_vid=01G1B1213YRBQAA1DA8467TWVA; adr_id=5fyNQvYT63eL24bCq2PNcuuOoLiZZsdEUPF1Md4bXq1SX1xj; login_ever=yes; first_visit_datetime=2022-04-26+22%3A14%3A36; webp_available=1; __utmv=235335808.|2=login%20ever=yes=1^3=plan=normal=1^5=gender=male=1^6=user_id=77510984=1^9=p_ab_id=7=1^10=p_ab_id_2=7=1^11=lang=zh=1^20=webp_available=yes=1; _gid=GA1.2.222557501.1651323352; _gcl_au=1.1.2062937937.1651323401; _td=9dd38fd1-01dd-4ed4-ae2f-aec8ac853f78; tag_view_ranking=_EOd7bsGyl~ziiAzr_h04~yREQ8PVGHN~0xsDLqCEW6~Ie2c51_4Sp~Lt-oEicbBr~sAwDH104z0~BC84tpS1K_~9Gbahmahac~eAg20hju2a~hW_oUTwHGx~3gc3uGrU1V~6qtAxDMeYD~1kqPgx5bT5~9ODMAZ0ebV~uW5495Nhg-~-98s6o2-Rp~2RR-Wztsl7~KMpT0re7Sq~hQnmcwWnXH~hRkZZnS6_e~RAtkGhC25o~iAHff6Sx6z~uC2yUZfXDc~aPdvNeJ_XM~28gdfFXlY7~TqiZfKmSCg~5oPIfUbtd6~9V46Zz_N_N~MDpxawlkUA~QB6vknyHRe~EZQqoW9r8g~2R7RYffVfj~r2kFZuczQC~6ib_FAsWf9~pnCQRVigpy~9TPuxVVpm_~KN7uxuR89w~VZAPYEMQnQ~eYfNGKUJaM~XhOHJMaDOw~59dAqNEUGJ~796B9WnMty~jYnWl04aAC~VIOKa7rioU~bzA-yjKTcQ~Bl1Zf-tv9I~MZZr-4xtm2~q303ip6Ui5~jImOpI7tih~MSNRmMUDgC~VN7cgWyMmg~6zMgdtdmwu~BSlt10mdnm~Kgk8e6lkPS~W7xIh3M3Xa~nPAsQaXukW~8_R-Kk5ORv~kbxsiTB3dV~yzNma8k7tF~DnqxiVoJax~LL4uU6tZ3S~8p7FrLtVHU~DNyuLA8MjA~pNxd7Oa9fU~HgUosxtViU~B_OtVkMSZT~qWFESUmfEs~q2ccYJV4vA~msnM7v3qDS~4tRDi-0MGG~0RlPEGkN4V~OeGEAgxZgp~1NmdBLOfGO~PsltMJiybA~2OuHYpR2h4~uBATiCdLrd~9yMCP-K2-R~vzTU7cI86f~RcRplUvn0G~GeMrue5H_4~E22Fw6ZrnB~xQBQg7AwXl~cAD2Df2Aaz~o3BmTZj2WP~TZdT4svccf~55eUPCA1fQ~ZJC17bcQNT~OH5ieNjgYI~LF9kqwfMs-~g0_q50TEM2~AoKfsFwwdu~8-iylRZUWo~bYxiZ03aun~4OHWT9rEGI~C-iDpIvyiK~jyw53VSia0~sFoUChFjE4~z_KWbpE2aS~txZ9z5ByU7; _im_uid.3929=b.81dfdb33a6e54e9b; __utma=235335808.1541353225.1650711335.1651384744.1651387311.15; __utmc=235335808; __utmz=235335808.1651387311.15.10.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmt=1; tags_sended=1; categorized_tags=4dC76RDj60~8NfvpmigcD~EZQqoW9r8g~IVwLyT8B6k~Ig5OcZugU6~hRkZZnS6_e~jYnWl04aAC; __cf_bm=5z.CVNE.srsZhJJjgek3SQqR3hYjy9lJeGScYn81l1g-1651387391-0-AacNU3yQgJCfm58zx8Dp/8j50dfvRkZJ5MFJo3MCxzT3XSGoV++VICo51JMZZJp394iQYmYybz/OKvBlAwIIFoOSdqcVUOgYeIy5obncRwrr0CNsySshHIX42hLNsmyLq9FYDZ5UDtAVjIMmLwIL5wZO7jrud2uqb2TfNaKf/UDFykVlVafsOxb/ilKbKxg5Eg==; PHPSESSID=77510984_xGmUhep1BaILFsq2MdGUOkbHGIic0iGn; _gat_UA-1830249-3=1; __utmb=235335808.3.10.1651387311; QSI_S_ZN_5hF4My7Ad6VNNAi=r:10:56; cto_bundle=KLUbn19FaVNyTHl3TGhRWlIlMkY1d01FY1pUSlJCWmlBY1JncjI0M0UzMFpYd2lBNWN2b0ZqTng2VFIxc3lMSGZPUEJ6NDlqa0VKZ08wdE5mYU9mdkFxM0ViYm5FWG96RExSNW10YlNQVVZMME8xZjhzc1pWYXBWUGQzbXlJakFaWGMzclFoMGR6dTNjSiUyRlYwTTR4NTNzY1ZVczdRJTNEJTNE"
	user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"
)


// url 
var (

	// everyday recommend page
	everydayRecommendPage = "https://www.pixiv.net/ranking.php?mode=daily_r18&content=illust"

	//proxy
	proxyUrl,_ =  url.Parse("http://127.0.0.1:5890")
)

// regex rule

var (

	// raw jpg rule
	// 
	rawJpgRule = `data-src="(?P<src>https://i.pximg.net/c/\d{1,}x\d{1,}/img-master/img/\d{4}/\d{2}/\d{2}/\d{2}/\d{2}/\d{2}/\d{8}_p0_master1200.jpg)"`
)

// local path
var (
	localPath = "/home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!.jpg"
)
