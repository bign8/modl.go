package unpacker

import (
	"bytes"
	"encoding/json"
	"testing"
)

var unpackerTests = []UnpackerTest{
	// Source: https://www.notion.so/num/Unpacker-Specification-f3385257766e44b4abf70fec4650ff7d
	{
		Name: "all",
		Input: `{
			"?": ["https://www.widgetcompany.com", "team", "janesmith", "johnwilson", "dashnaanand"],
			"o": [
			  "Widget Company Ltd",
			  "Making the best widgets",
			  "%0",
			  [
				["Jane Smith", "%ceo", "%0%/%1%/%2", "%2", "%2%widgets"],
				["John Wilson", "%cto", "%0%/%1%/%3", "%3", "jono"],
				["Dashna Anand", "%cmo", "%0%/%1%/%4", "%4", "%4"]
			  ]
			]
		  }`,
		Subs: `{
			"ceo": "Chief Executive Officer",
			"cto": "Chief Technology Officer",
			"cmo": "Chief Marketing Officer"
		  }`,
		Trans: `{
			"o": {
			  "assignKeys": ["n", "s", "w", "e"],
			  "replacePair": {
				"name": "%n",
				"strapline": "%s",
				"website": "%w",
				"employees": "%e"
			  }
			},
			"e" : {
			  "arrayItems" : "em"
			},
			"em": {
			  "assignKeys": ["n", "p", "b", "l", "t"],
			  "replacePair": {
				"name": "%n",
				"position": "%p",
				"bio": "%b",
				"linkedin": "https://www.linkedin.com/in/%l",
				"twitter": "https://www.twitter.com/%t"
			  }
			}
		  }`,
		Output: `{
			"name": "Widget Company Ltd",
			"strapline": "Making the best widgets",
			"website": "https://www.widgetcompany.com",
			"employees": [{
			  "name": "Jane Smith",
			  "position": "Chief Executive Officer",
			  "bio": "https://www.widgetcompany.com/team/janesmith",
			  "linkedin": "https://www.linkedin.com/in/janesmith",
			  "twitter": "https://www.twitter.com/janesmithwidgets"
			}, {
			  "name": "John Wilson",
			  "position": "Chief Technology Officer",
			  "bio": "https://www.widgetcompany.com/team/johnwilson",
			  "linkedin": "https://www.linkedin.com/in/johnwilson",
			  "twitter": "https://www.twitter.com/jono"
			}, {
			  "name": "Dashna Anand",
			  "position": "Chief Marketing Officer",
			  "bio": "https://www.widgetcompany.com/team/dashnaanand",
			  "linkedin": "https://www.linkedin.com/in/dashnaanand",
			  "twitter": "https://www.twitter.com/dashnaanand"
			}]
		  }`,
	}, {
		Name:   "rewriteKey",
		Input:  `{"f": 1}`,
		Trans:  `{"f": {"rewriteKey": "foo"}}`,
		Output: `{"foo": 1}`,
	}, {
		Name:   "rewriteValue",
		Input:  `{"foo": 1}`,
		Trans:  `{"foo": {"rewriteValue" : { "x" : "%self", "y" : 2 }}}`,
		Output: `{"foo": {"x": 1, "y": 2}}`,
	}, {
		Name:   "replacePair",
		Input:  `{"f": 1}`,
		Trans:  `{"f": {"replacePair": {"x": "%self", "y": 2, "z": 3}}}`,
		Output: `{"x" : 1, "y" : 2, "z" : 3}`,
	}, {
		Name:   "subs",
		Input:  `{"foo": "%var"}`,
		Subs:   `{"var": 1}`,
		Output: `{"foo": 1}`,
	}, {
		Name:   "assignKeys",
		Input:  `{"phonetic": ["alpha", "bravo", "charlie"]}`,
		Trans:  `{"phonetic": {"assignKeys": ["a", "b", "c"]}}`,
		Output: `{"phonetic": {"a": "alpha", "b": "bravo", "c": "charlie"}}`,
	}, {
		Name:  "arrayItems",
		Input: `{"employees": [["Jane Smith", "CEO"], ["Dashna Anand", "CMO"]]}`,
		Trans: `{
			"employees": {"arrayItems": "employee"},
			"employee": {
				"assignKeys": ["name", "position"],
				"replacePair": {"name": "%name", "position": "%position"}
			}}`,
		Output: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
	}, {
		Name: "noop",
		Input: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
		Output: `{
			"employees": [
				{"name": "Jane Smith", "position": "CEO"},
				{"name": "Dashna Anand", "position": "CMO"}
			]}`,
	},

	// Source: https://www.unpacker.uk/playground?e=3
	{
		Name: "playground-3",
		Skip: true,
		Input: `{
			"@n": 1,
			"o": {
			  "n": "NUM Example Co",
			  "s": "Example Strapline",
			  "c": [
				{
				  "t": {
					"d": "Customer Service",
					"v": "+441270123456"
				  }
				},
				{
				  "fb": {
					"v": "examplefacebook"
				  }
				},
				{
				  "in": {
					"v": "exampleinstagram"
				  }
				},
				{
				  "tw": {
					"v": "exampletwitter"
				  }
				}
			  ]
			}
		  }`,
		Subs: `{
			"locale": {
			  "p": {
				"name": "Person"
			  },
			  "gr": {
				"name": "Group"
			  },
			  "o": {
				"name": "Organisation"
			  },
			  "dp": {
				"name": "Department"
			  },
			  "e": {
				"name": "Employee"
			  },
			  "lc": {
				"name": "Location"
			  },
			  "gp": {
				"name": "Group"
			  },
			  "t": {
				"name": "Telephone",
				"default": "Call"
			  },
			  "sm": {
				"name": "SMS",
				"default": "Text"
			  },
			  "u": {
				"name": "Web URL",
				"default": "Click"
			  },
			  "uu": {
				"name": "Web URL (http - unsecure)",
				"default": "Click"
			  },
			  "g": {
				"name": "GPS",
				"default": "View Location"
			  },
			  "a": {
				"name": "Address",
				"default": "View Address"
			  },
			  "fx": {
				"name": "Fax",
				"default": "Send a fax"
			  },
			  "em": {
				"name": "Email",
				"default": "Send an email"
			  },
			  "aa": {
				"name": "Android App",
				"default": "Download the app"
			  },
			  "as": {
				"name": "iOS App",
				"default": "Download the app"
			  },
			  "bt": {
				"name": "Baidu Tieba",
				"default": "View Baidu profile"
			  },
			  "fb": {
				"name": "Facebook",
				"default": "View Facebook profile"
			  },
			  "fs": {
				"name": "FourSquare",
				"default": "View FourSquare page"
			  },
			  "ft": {
				"name": "FaceTime",
				"default": "Call with Facetime"
			  },
			  "gh": {
				"name": "Github",
				"default": "View Github profile"
			  },
			  "im": {
				"name": "iMessage",
				"default": "Send iMessage"
			  },
			  "in": {
				"name": "Instagram",
				"default": "View Instagram profile"
			  },
			  "kk": {
				"name": "Kik",
				"default": "Connect with Kik"
			  },
			  "li": {
				"name": "LinkedIn",
				"default": "View LinkedIn page"
			  },
			  "ln": {
				"name": "Line",
				"default": "Connect with Line"
			  },
			  "md": {
				"name": "Medium",
				"default": "View Medium blog"
			  },
			  "pr": {
				"name": "Periscope",
				"default": "View Periscope profile"
			  },
			  "pi": {
				"name": "Pinterest",
				"default": "View Pinterest board"
			  },
			  "qq": {
				"name": "QQ",
				"default": "View QQ Page"
			  },
			  "qz": {
				"name": "Qzone",
				"default": "View Qzone Page"
			  },
			  "rd": {
				"name": "Reddit",
				"default": "View subreddit"
			  },
			  "rn": {
				"name": "Renren",
				"default": "View Renren profile"
			  },
			  "sc": {
				"name": "Soundcloud",
				"default": "View Soundcloud page"
			  },
			  "sk": {
				"name": "Skype",
				"default": "Call with Skype"
			  },
			  "sr": {
				"name": "Swarm",
				"default": "Connect with Swarm"
			  },
			  "sn": {
				"name": "Snapchat",
				"default": "Connect with Snapchat"
			  },
			  "sw": {
				"name": "Sina Weibo",
				"default": "View Weibo page"
			  },
			  "tb": {
				"name": "Tumblr",
				"default": "View Tumblr blog"
			  },
			  "tl": {
				"name": "Telegram",
				"default": "Connect with Telegram"
			  },
			  "tw": {
				"name": "Twitter",
				"default": "View Twitter profile"
			  },
			  "to": {
				"name": "Twoo",
				"default": "View Twoo page"
			  },
			  "vb": {
				"name": "Viber",
				"default": "Call with Viber"
			  },
			  "vk": {
				"name": "Vkontakte",
				"default": "View VK page"
			  },
			  "vm": {
				"name": "Vimeo",
				"default": "View Vimeo profile"
			  },
			  "wa": {
				"name": "Whatsapp",
				"default": "Message on Whatsapp"
			  },
			  "wc": {
				"name": "WeChat",
				"default": "Connect with WeChat"
			  },
			  "xi": {
				"name": "Xing",
				"default": "View Xing page"
			  },
			  "yt": {
				"name": "YouTube",
				"default": "View YouTube channel"
			  },
			  "yy": {
				"name": "YY",
				"default": "View YY page"
			  }
			},
			"AC": "Accounts",
			"CS": "Customer Service"
		  }`,
		Trans: `{
			"o": {
			  "rewriteKey": "organisation",
			  "assignKeys": [
				"n",
				"s",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.o.name",
				"name": "%n",
				"slogan": "%s",
				"contacts": "%c"
			  }
			},
			"p": {
			  "rewriteKey": "person",
			  "assignKeys": [
				"n",
				"b",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.p.name",
				"name": "%n",
				"bio": "%b",
				"contacts": "%c"
			  }
			},
			"e": {
			  "rewriteKey": "employee",
			  "assignKeys": [
				"n",
				"r",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.e.name",
				"name": "%n",
				"role": "%r",
				"contacts": "%c"
			  }
			},
			"lc": {
			  "rewriteKey": "location",
			  "assignKeys": [
				"n",
				"d",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.lc.name",
				"name": "%n",
				"description": "%d",
				"contacts": "%c"
			  }
			},
			"gp": {
			  "rewriteKey": "group",
			  "assignKeys": [
				"n",
				"d",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.gp.name",
				"name": "%n",
				"description": "%d",
				"contacts": "%c"
			  }
			},
			"dp": {
			  "rewriteKey": "department",
			  "assignKeys": [
				"n",
				"d",
				"c"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.dp.name",
				"name": "%n",
				"description": "%d",
				"contacts": "%c"
			  }
			},
			"a": {
			  "rewriteKey": "address",
			  "assignKeys": [
				"al",
				"pz",
				"co",
				"d"
			  ],
			  "rewriteValue": {
				"description": "%d",
				"description_default": "%/subs.locale.a.default",
				"lines": "%al",
				"postcode": "%pz",
				"country": "%co",
				"method_type": "core",
				"object_display_name": "%/subs.locale.a.name",
				"prefix": ""
			  }
			},
			"l": {
			  "rewriteKey": "link",
			  "assignKeys": [
				"@L",
				"d"
			  ],
			  "rewriteValue": {
				"@L": "%@L",
				"description": "%d"
			  }
			},
			"fb": {
			  "rewriteKey": "facebook",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.fb.name",
				"description_default": "%/subs.locale.fb.default",
				"description": "%d",
				"prefix": "https://www.facebook.com/",
				"method_type": "third_party",
				"controller": "facebook.com",
				"value": "%v"
			  }
			},
			"g": {
			  "rewriteKey": "gps",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.g.name",
				"description_default": "%/subs.locale.g.default",
				"description": "%d",
				"prefix": "",
				"method_type": "core",
				"value": "%v"
			  }
			},
			"in": {
			  "rewriteKey": "instagram",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.in.name",
				"description_default": "%/subs.locale.in.default",
				"description": "%d",
				"prefix": "https://www.instagram.com/",
				"method_type": "third_party",
				"controller": "instagram.com",
				"value": "%v"
			  }
			},
			"li": {
			  "rewriteKey": "linkedin",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.li.name",
				"description_default": "%/subs.locale.li.default",
				"description": "%d",
				"prefix": "https://www.linkedin.com/",
				"method_type": "third_party",
				"controller": "linkedin.com",
				"value": "%v"
			  }
			},
			"yt": {
			  "rewriteKey": "youtube",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.yt.name",
				"description_default": "%/subs.locale.yt.default",
				"description": "%d",
				"prefix": "https://www.youtube.com/",
				"method_type": "third_party",
				"controller": "youtube.com",
				"value": "%v"
			  }
			},
			"pi": {
			  "rewriteKey": "pinterest",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.pi.name",
				"description_default": "%/subs.locale.pi.default",
				"description": "%d",
				"prefix": "https://www.pinterest.com/",
				"method_type": "third_party",
				"controller": "pinterest.com",
				"value": "%v"
			  }
			},
			"tw": {
			  "rewriteKey": "twitter",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.tw.name",
				"description_default": "%/subs.locale.tw.default",
				"description": "%d",
				"prefix": "https://www.twitter.com/",
				"method_type": "third_party",
				"controller": "twitter.com",
				"value": "%v",
				"value_prefix": "@"
			  }
			},
			"t": {
			  "rewriteKey": "telephone",
			  "assignKeys": [
				"v",
				"d",
				"h"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.t.name",
				"description_default": "%/subs.locale.t.default",
				"description": "%d",
				"prefix": "tel:",
				"method_type": "core",
				"value": "%v",
				"hours": "%h"
			  }
			},
			"sm": {
			  "rewriteKey": "sms",
			  "assignKeys": [
				"v",
				"d",
				"h"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.sm.name",
				"description_default": "%/subs.locale.sm.default",
				"description": "%d",
				"prefix": "sms:",
				"method_type": "core",
				"value": "%v",
				"hours": "%h"
			  }
			},
			"em": {
			  "rewriteKey": "email",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.em.name",
				"description_default": "%/subs.locale.em.default",
				"description": "%d",
				"prefix": "mailto:",
				"method_type": "core",
				"value": "%v"
			  }
			},
			"fx": {
			  "rewriteKey": "fax",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.fx.name",
				"description_default": "%/subs.locale.fx.default",
				"description": "%d",
				"prefix": "tel:",
				"method_type": "core",
				"value": "%v"
			  }
			},
			"u": {
			  "rewriteKey": "url",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.u.name",
				"description_default": "%/subs.locale.u.default",
				"description": "%d",
				"prefix": "https://",
				"method_type": "core",
				"value": "%v"
			  }
			},
			"uu": {
			  "rewriteKey": "unsecure_url",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "%/subs.locale.uu.name",
				"description_default": "%/subs.locale.uu.default",
				"description": "%d",
				"prefix": "http://",
				"method_type": "core",
				"value": "%v"
			  }
			},
			"av": {
			  "rewriteKey": "available"
			},
			"tz": {
			  "rewriteKey": "time_zone_location"
			},
			"i": {
			  "rewriteKey": "introduction"
			},
			"ac": {
			  "rewriteKey": "access"
			},
			"aa": {
			  "rewriteKey": "android-app",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.aa.name",
				"description_default": "locale.aa.default",
				"prefix": "https://play.google.com/store/apps/details?id=",
				"method_type": "third_party",
				"controller": "play.google.com"
			  }
			},
			"as": {
			  "rewriteKey": "ios-app",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.as.name",
				"description_default": "locale.as.default",
				"prefix": "https://itunes.apple.com/app/",
				"method_type": "third_party",
				"controller": "apps.apple.com"
			  }
			},
			"bt": {
			  "rewriteKey": "baidu_tieba",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.bt.name",
				"description_default": "locale.bt.default",
				"prefix": "https://tieba.baidu.com/",
				"method_type": "third_party",
				"controller": "tieba.baidu.com"
			  }
			},
			"fs": {
			  "rewriteKey": "foursquare",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.fs.name",
				"description_default": "locale.fs.default",
				"prefix": "https://www.foursquare.com/",
				"method_type": "third_party",
				"controller": "foursquare.com"
			  }
			},
			"ft": {
			  "rewriteKey": "facetime",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.ft.name",
				"description_default": "locale.ft.default",
				"prefix": "facetime://",
				"method_type": "third_party",
				"controller": "facetime@apple.com"
			  }
			},
			"gh": {
			  "rewriteKey": "github",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.gh.name",
				"description_default": "locale.gh.default",
				"prefix": "https://www.github.com/",
				"method_type": "third_party",
				"controller": "github.com"
			  }
			},
			"im": {
			  "rewriteKey": "imessage",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.im.name",
				"description_default": "locale.im.default",
				"prefix": "imessage://",
				"method_type": "third_party",
				"controller": "imessage@apple.com"
			  }
			},
			"kk": {
			  "rewriteKey": "kik",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.kk.name",
				"description_default": "locale.kk.default",
				"prefix": "https://www.kik.com/u/",
				"method_type": "third_party",
				"controller": "kik.com"
			  }
			},
			"ln": {
			  "rewriteKey": "line",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.ln.name",
				"description_default": "locale.ln.default",
				"prefix": "line://",
				"method_type": "third_party",
				"controller": "line.me"
			  }
			},
			"md": {
			  "rewriteKey": "medium",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.md.name",
				"description_default": "locale.md.default",
				"prefix": "https://www.medium.com/",
				"method_type": "third_party",
				"controller": "medium.com"
			  }
			},
			"pr": {
			  "rewriteKey": "periscope",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.pr.name",
				"description_default": "locale.pr.default",
				"prefix": "https://www.periscope.tv/",
				"method_type": "third_party",
				"controller": "periscope.tv"
			  }
			},
			"qq": {
			  "rewriteKey": "qq",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.qq.name",
				"description_default": "locale.qq.default",
				"prefix": "https://www.qq.com/",
				"method_type": "third_party",
				"controller": "qq.com"
			  }
			},
			"qz": {
			  "rewriteKey": "qzone",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.qz.name",
				"description_default": "locale.qz.default",
				"prefix": "https://www.qzone.com/",
				"method_type": "third_party",
				"controller": "qzone.com"
			  }
			},
			"rd": {
			  "rewriteKey": "reddit",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.rd.name",
				"description_default": "locale.rd.default",
				"prefix": "https://www.reddit.com/r/",
				"method_type": "third_party",
				"controller": "reddit.com"
			  }
			},
			"rn": {
			  "rewriteKey": "renren",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.rn.name",
				"description_default": "locale.rn.default",
				"prefix": "https://www.renren.com/",
				"method_type": "third_party",
				"controller": "renren.com"
			  }
			},
			"sc": {
			  "rewriteKey": "soundcloud",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.sc.name",
				"description_default": "locale.sc.default",
				"prefix": "https://www.soundcloud.com/",
				"method_type": "third_party",
				"controller": "soundcloud.com"
			  }
			},
			"sk": {
			  "rewriteKey": "skype",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.sk.name",
				"description_default": "locale.sk.default",
				"prefix": "skype:",
				"method_type": "third_party",
				"controller": "skype.com"
			  }
			},
			"sr": {
			  "rewriteKey": "swarm",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.sr.name",
				"description_default": "locale.sr.default",
				"prefix": "https://www.swarmapp.com/",
				"method_type": "third_party",
				"controller": "swarmapp.com"
			  }
			},
			"sn": {
			  "rewriteKey": "snapchat",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.sn.name",
				"description_default": "locale.sn.default",
				"prefix": "snapchat://add/",
				"method_type": "third_party",
				"controller": "snapchat.com"
			  }
			},
			"sw": {
			  "rewriteKey": "sina-weibo",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.sw.name",
				"description_default": "locale.sw.default",
				"prefix": "https://www.weibo.com/",
				"method_type": "third_party",
				"controller": "weibo.com"
			  }
			},
			"tb": {
			  "rewriteKey": "tumblr",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.tb.name",
				"description_default": "locale.tb.default",
				"prefix": "https://.tumblr.com/",
				"method_type": "third_party",
				"controller": "tumblr.com"
			  }
			},
			"tl": {
			  "rewriteKey": "telegram",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.tl.name",
				"description_default": "locale.tl.default",
				"prefix": "https://www.telegram.me/",
				"method_type": "third_party",
				"controller": "telegram.com"
			  }
			},
			"to": {
			  "rewriteKey": "twoo",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.to.name",
				"description_default": "locale.to.default",
				"prefix": "https://www.twoo.com/",
				"method_type": "third_party",
				"controller": "twoo.com"
			  }
			},
			"vb": {
			  "rewriteKey": "viber",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.vb.name",
				"description_default": "locale.vb.default",
				"prefix": "https://www.viber.com/",
				"method_type": "third_party",
				"controller": "viber.com"
			  }
			},
			"vk": {
			  "rewriteKey": "vkontakte",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.vk.name",
				"description_default": "locale.vk.default",
				"prefix": "https://www.vk.com/",
				"method_type": "third_party",
				"controller": "vk.com"
			  }
			},
			"vm": {
			  "rewriteKey": "vimeo",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.vm.name",
				"description_default": "locale.vm.default",
				"prefix": "https://www.vimeo.com/",
				"method_type": "third_party",
				"controller": "vimeo.com"
			  }
			},
			"wa": {
			  "rewriteKey": "whatsapp",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.wa.name",
				"description_default": "locale.wa.default",
				"prefix": "whatsapp://",
				"method_type": "third_party",
				"controller": "whatsapp.com"
			  }
			},
			"wc": {
			  "rewriteKey": "wechat",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.wc.name",
				"description_default": "locale.wc.default",
				"prefix": "https://www.wechat.com/",
				"method_type": "third_party",
				"controller": "wechat.com"
			  }
			},
			"xi": {
			  "rewriteKey": "xing",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.xi.name",
				"description_default": "locale.xi.default",
				"prefix": "https://www.xing.com/",
				"method_type": "third_party",
				"controller": "xing.com"
			  }
			},
			"yy": {
			  "rewriteKey": "yy",
			  "assignKeys": [
				"v",
				"d"
			  ],
			  "rewriteValue": {
				"object_display_name": "locale.yy.name",
				"description_default": "locale.yy.default",
				"prefix": "https://www.yy.com/",
				"method_type": "third_party",
				"controller": "yy.com"
			  }
			}
		  }`,
		Output: `{
			"@n": 1,
			"organisation": {
			  "object_display_name": "Organisation",
			  "name": "NUM Example Co",
			  "slogan": "Example Strapline",
			  "contacts": [
				{
				  "telephone": {
					"object_display_name": "Telephone",
					"description_default": "Call",
					"description": "Customer Service",
					"prefix": "tel:",
					"method_type": "core",
					"value": "+441270123456",
					"hours": null
				  }
				},
				{
				  "facebook": {
					"object_display_name": "Facebook",
					"description_default": "View Facebook profile",
					"description": null,
					"prefix": "https://www.facebook.com/",
					"method_type": "third_party",
					"controller": "facebook.com",
					"value": "examplefacebook"
				  }
				},
				{
				  "instagram": {
					"object_display_name": "Instagram",
					"description_default": "View Instagram profile",
					"description": null,
					"prefix": "https://www.instagram.com/",
					"method_type": "third_party",
					"controller": "instagram.com",
					"value": "exampleinstagram"
				  }
				},
				{
				  "twitter": {
					"object_display_name": "Twitter",
					"description_default": "View Twitter profile",
					"description": null,
					"prefix": "https://www.twitter.com/",
					"method_type": "third_party",
					"controller": "twitter.com",
					"value": "exampletwitter",
					"value_prefix": "@"
				  }
				}
			  ]
			}
		  }`,
	},

	// Source: https://www.unpacker.uk/playground?e=2
	{
		Name: "playground-2",
		Input: `{
			"?": ["https://www.widgetcompany.com", "team", "janesmith", "johnwilson", "dashnaanand"],
			"o": [
			  "Widget Company Ltd",
			  "Making the best widgets",
			  "%0",
			  [
				["Jane Smith", "%ceo", "%0%/%1%/%2", "%2", "%2%widgets"],
				["John Wilson", "%cto", "%0%/%1%/%3", "%3", "jono"],
				["Dashna Anand", "%cmo", "%0%/%1%/%4", "%4", "%4"]
			  ]
			]
		  }`,
		Subs: `{
			"ceo": "Chief Executive Officer",
			"cto": "Chief Technology Officer",
			"cmo": "Chief Marketing Officer"
		  }`,
		Trans: `{
			"o": {
			  "assignKeys": ["n", "s", "w", "e"],
			  "replacePair": {
				"name": "%n",
				"strapline": "%s",
				"website": "%w",
				"employees": "%e"
			  }
			},
			"em": {
			  "replacePair": {
				"name": "%n",
				"position": "%p",
				"bio": "%b",
				"linkedin": "https://www.linkedin.com/in/%l",
				"twitter": "https://www.twitter.com/%t"
			  }
			}
		  }`,
		Output: `{
			"name": "Widget Company Ltd",
			"strapline": "Making the best widgets",
			"website": "https://www.widgetcompany.com",
			"employees": [
			  [
				"Jane Smith",
				"Chief Executive Officer",
				"https://www.widgetcompany.com/team/janesmith",
				"janesmith",
				"janesmithwidgets"
			  ],
			  [
				"John Wilson",
				"Chief Technology Officer",
				"https://www.widgetcompany.com/team/johnwilson",
				"johnwilson",
				"jono"
			  ],
			  [
				"Dashna Anand",
				"Chief Marketing Officer",
				"https://www.widgetcompany.com/team/dashnaanand",
				"dashnaanand",
				"dashnaanand"
			  ]
			]
		  }`,
	},

	// Source: https://www.unpacker.uk/playground?e=1
	{
		Name: "playground-1",
		Input: `{
			"?": ["abccompany", "+4412345678"],
			"c": "ABC Company Ltd",
			"t": [
			  {"l": "%ac", "n": "%1%90"},
			  {"l": "%cs", "n": "%1%89"}
			],
			"tw": "/%0",
			"i": "/%0%pics"
		  }`,
		Subs: `{"cs": "Customer Service", "ac": "Accounts"}`,
		Trans: `{
			"c": {"rewriteKey": "coname"},
			"t": {"rewriteKey": "telephone"},
			"l": {"rewriteKey": "label"},
			"n": {"rewriteKey": "number"},
			"tw": {"rewriteKey": "twitter"},
			"i": {"rewriteKey": "instagram"}
		  }`,
		Output: `{
			"coname": "ABC Company Ltd",
			"telephone": [
			  {"label": "Accounts", "number": "+441234567890"},
			  {"label": "Customer Service", "number": "+441234567889"}
			],
			"twitter": "/abccompany",
			"instagram": "/abccompanypics"
		  }`,
	},

	// Source: https://www.unpacker.uk/specification#variable-index
	{
		Name:   "variable-index-1",
		Input:  `{"?": [1], "a": "%0%"}`,
		Output: `{"a": 1}`,
	},
	{
		Name:   "variable-index-2",
		Input:  `{"?": [["a", "b"], {"a": "alpha", "b": "bravo"}], "x": "%0%", "y": "%1%"}`,
		Output: `{"x": ["a", "b"], "y": {"a": "alpha", "b": "bravo"}}`,
	},
	{
		Name:   "variable-index-3",
		Input:  `{"?": [["a", "b"], {"a": "alpha", "b": "bravo"}], "x": "%0.0%", "y": "%1.a%"}`,
		Output: `{"x": "a", "y": "alpha"}`,
	},

	// Source: https://www.unpacker.uk/specification#substitution-object
	{
		Name:   "sub-1",
		Input:  `{"this": "%a%", "that": "%b%"}`,
		Subs:   `{"a": 1, "b": 2}`,
		Output: `{"this": 1, "that": 2}`,
	},
	{
		Name:   "sub-2",
		Input:  `{"this": "%a.x%", "that": "%b.0%"}`,
		Subs:   `{"a": {"x": "xray", "y": "yankee"}, "b": [1, 2]}`,
		Output: `{"this": "xray", "that": 1}`,
	},

	// Source: https://www.unpacker.uk/specification#transformation-object
	{
		Name: "trans-1",
		Input: `{
			"t": 1,
			"u": {"t": 2},
			"v": {"w": {"t": 3}}
		  }`,
		Trans: `{"t": {"rewriteKey": "testing"}}`,
		Output: `{
			"testing": 1,
			"u": {"testing": 2},
			"v": {"w": {"testing": 3}}
		  }`,
	},
	{
		Name: "trans-2",
		Skip: true,
		Input: `{
			"t": 1,
			"u": {
			  "t": 2
			},
			"v": {
			  "w": {
				"t": 3
			  }
			}
		  }`,
		Trans: `{
			"t": {
			  "rewriteKey": "testing"
			},
			"u": {
			  "t": {
				"rewriteKey": "testing2"
			  }
			},
			"v": {
			  "w": {
				"t": {
				  "rewriteKey": "testing3"
				}
			  }
			}
		  }`,
		Output: `{
			"testing": 1,
			"u": {
			  "testing2": 2
			},
			"v": {
			  "w": {
				"testing3": 3
			  }
			}
		  }`,
	},
	{
		Name:   "trans-3",
		Input:  `{"t": 1}`,
		Trans:  `{"t": {"rewriteKey": "testing"}}`,
		Output: `{"testing": 1}`,
	},
	{
		Name:  "trans-4",
		Input: `{"t": "me"}`,
		Trans: `{
			"t": {
			  "rewriteKey": "twitter",
			  "rewriteValue": "twitter.com/%self"
			}
		  }`,
		Output: `{"twitter": "twitter.com/me"}`,
	},
	{
		Name:  "trans-5",
		Input: `{"t": "me"}`,
		Trans: `{
			"t": {
			  "rewriteKey": "twitter",
			  "rewriteValue": {
				"cta": "Follow on Twitter",
				"url": "twitter.com/%self"
			  }
			}
		  }`,
		Output: `{
			"twitter": {
			  "cta": "Follow on Twitter",
			  "url": "twitter.com/me"
			}
		  }`,
	},
	{
		Name:  "trans-6",
		Input: `{"t": "me"}`,
		Trans: `{
			"t": {
			  "replacePair": {
				"cta": "Follow on Twitter",
				"url": "twitter.com/%self"
			  }
			}
		  }`,
		Output: `{"cta": "Follow on Twitter", "url": "twitter.com/me"}`,
	},
	{
		Name:  "trans-7",
		Input: `{"phonetic": ["alpha", "bravo", "charlie"]}`,
		Trans: `{"phonetic": {"assignKeys": ["a", "b", "c"]}}`,
		Output: `{
			"phonetic": {
			  "a": "alpha",
			  "b": "bravo",
			  "c": "charlie"
			}
		  }`,
	},
	{
		Name: "trans-8",
		Input: `{
			"employees": [
			  ["Jane Smith", "CEO"],
			  ["Dashna Anand", "CMO"]
			]
		  }`,
		Trans: `{
			"employees": {
			  "arrayItems": "employee"
			},
			"employee": {
			  "assignKeys": [
				"name",
				"position"
			  ],
			  "replacePair": {
				"name": "%name",
				"position": "%position"
			  }
			}
		  }`,
		Output: `{"employees": [
			  {"name": "Jane Smith", "position": "CEO"},
			  {"name": "Dashna Anand", "position": "CMO"}
			]}`,
	},

	// Source: https://www.unpacker.uk/specification#referencing
	{
		Name: "ref-1",
		Input: `{
			"?": ["test"],
			"a": "%0%",
			"b": "this is a %0",
			"c": "this is another %0 of referencing",
			"d": "once again %0%ing it"
		  }`,
		Output: `{
			"a": "test",
			"b": "this is a test",
			"c": "this is another test of referencing",
			"d": "once again testing it"
		  }`,
	},
	{
		Name: "ref-2",
		Input: `{
			"?": [
			  ["a", "b"],
			  {"a": "alpha", "b": "bravo"}
			],
			"x": "%0.0",
			"y": "%1.a",
			"z": "%letters.0"
		  }`,
		Subs:   `{"letters": ["a", "b"]}`,
		Output: `{"x": "a", "y": "alpha", "z": "a"}`,
	},
	{
		Name:   "ref-3",
		Input:  `{"a": {"b": {"c": 1}}}`,
		Trans:  `{"a": {"rewriteValue": "%b.c"}}`,
		Output: `{"a": 1}`,
	},
	{
		Name: "ref-4",
		Input: `{
			"a": ["alpha", "bravo", "charlie"]
		  }`,
		Trans: `{
			"a": {
			  "assignKeys": ["a", "b", "c"],
			  "rewriteValue": "%c"
			}
		  }`,
		Output: `{"a": "charlie"}`,
	},
	{
		Name:  "ref-5",
		Input: `{"a": 1, "b": 2, "c": 3}`,
		Trans: `{
			"a": {"rewriteValue": "%/compact.c"},
			"b": {"rewriteValue": "%/subs.x"}
		  }`,
		Subs:   `{"x": 4}`,
		Output: `{"a": 3, "b": 4, "c": 3}`,
	},
}

type UnpackerTest struct {
	Name   string
	Input  string
	Subs   string
	Trans  string
	Output string
	Skip   bool
}

func (test *UnpackerTest) Run(t *testing.T) {
	if testing.Short() && test.Skip {
		t.Skip("manually disabled test")
	}
	u := &Unpacker{}
	if test.Trans != `` {
		trans, err := ParseTransforms([]byte(test.Trans))
		if err != nil {
			t.Fatalf("Unable to parse transforms: %v", err)
		}
		u.Transforms = trans
	}
	if test.Subs != `` {
		subs := make(Substitution)
		if err := json.Unmarshal([]byte(test.Subs), &subs); err != nil {
			t.Fatalf("Unable to parse substition: %v", err)
		}
		u.AddSubs(subs)
	}
	output, err := u.Unpack([]byte(test.Input))
	if err != nil {
		t.Fatalf("Unable to unpack: %v", err)
	}
	exp := test.out()
	if !bytes.Equal(output, exp) {
		t.Fatalf("Unexpected output:\n\tWant: %s\n\tGot:  %s", string(exp), string(output))
	}
}

func (test UnpackerTest) out() []byte {
	var obj interface{}
	err := json.Unmarshal([]byte(test.Output), &obj)
	if err != nil {
		panic(err)
	}
	bits, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return bits
}

func TestUnpacker(t *testing.T) {
	for _, test := range unpackerTests {
		t.Run(test.Name, test.Run)
	}
}
