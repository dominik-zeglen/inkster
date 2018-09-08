import { Theme } from "react-jss";

export default (theme: Theme) =>
  ({
    "@import": [
      "url('https://fonts.googleapis.com/css?family=Open+Sans:400,400i,600,600i')"
    ],
    "@global": {
      html: {
        ...theme.typography.body,
        color: theme.colors.black.main
      },
      body: {
        backgroundColor: theme.colors.white.main,
        margin: "0",
        transitionTimingFunction: 'ease-in' as 'ease-in'
      },
      "article, aside, details, figcaption, figure, footer, header, hgroup, main, menu, nav, section, summary": {
        display: "block"
      },
      "audio, canvas, progress, video": {
        display: "inline-block",
        verticalAlign: "baseline"
      },
      "audio:not([controls])": {
        display: "none",
        height: "0"
      },
      "[hidden], template": {
        display: "none"
      },
      a: {
        backgroundColor: "transparent",
        color: theme.colors.primary.main,
        textDecoration: "none"
      },
      "a:active, a:hover": {
        outline: "0"
      },
      "abbr[title]": {
        borderBottom: "1px dotted"
      },
      "b, strong": {
        fontWeight: "bold"
      },
      dfn: {
        fontStyle: "italic"
      },
      h1: {
        margin: ".67em 0",
        fontSize: "2em"
      },
      mark: {
        color: "#000",
        background: "#ff0"
      },
      small: {
        fontSize: "80%"
      },
      "sub, sup": {
        position: "relative",
        fontSize: "75%",
        lineHeight: "0",
        verticalAlign: "baseline"
      },
      sup: {
        top: "-.5em"
      },
      sub: {
        bottom: "-.25em"
      },
      img: {
        border: "0",
        verticalAlign: "middle"
      },
      "svg:not(:root)": {
        overflow: "hidden"
      },
      figure: {
        margin: "0",
        fallbacks: [
          {
            margin: "1em 40px"
          }
        ]
      },
      hr: {
        height: "0",
        W: "content-box",
        M: "content-box",
        boxSizing: "content-box",
        marginTop: 20,
        marginBottom: 20,
        border: "0",
        borderTop: "1px solid #eee"
      },
      pre: {
        overflow: "auto",
        display: "block",
        padding: 9.5,
        margin: "0 0 10px",
        fontSize: 13,
        lineHeight: "1.42857143",
        color: "#333",
        wordBreak: "break-all",
        wordWrap: "break-word",
        backgroundColor: "#f5f5f5",
        border: "1px solid #ccc",
        borderRadius: 4
      },
      "code, kbd, pre, samp": {
        fontFamily: 'Menlo, Monaco, Consolas, "Courier New", monospace',
        fontSize: "1em",
        fallbacks: [
          {
            fontFamily: "monospace, monospace"
          }
        ]
      },
      "button, input, optgroup, select, textarea": {
        margin: "0",
        font: "inherit",
        color: "inherit"
      },
      button: {
        overflow: "visible"
      },
      "button, select": {
        textTransform: "none"
      },
      'button, html input[type="button"], input[type="reset"], input[type="submit"]': {
        W: "button",
        cursor: "pointer"
      },
      "button[disabled], html input[disabled]": {
        cursor: "default"
      },
      "button::-moz-focus-inner, input::-moz-focus-inner": {
        padding: "0",
        border: "0"
      },
      input: {
        lineHeight: "normal"
      },
      'input[type="checkbox"], input[type="radio"]': {
        W: "border-box",
        M: "border-box",
        boxSizing: "border-box",
        padding: "0"
      },
      'input[type="number"]::-webkit-inner-spin-button, input[type="number"]::-webkit-outer-spin-button': {
        height: "auto"
      },
      'input[type="search"]': {
        W: "none",
        M: "border-box",
        boxSizing: "border-box",
        fallbacks: [
          {
            W: "border-box"
          },
          {
            boxSizing: "content-box"
          },
          {
            M: "content-box"
          },
          {
            W: "textfield"
          },
          {
            W: "content-box"
          }
        ]
      },
      'input[type="search"]::-webkit-search-cancel-button, input[type="search"]::-webkit-search-decoration': {
        W: "none"
      },
      fieldset: {
        padding: "0",
        margin: "0",
        border: "0",
        minWidth: "0",
        fallbacks: [
          {
            border: "1px solid #c0c0c0"
          },
          {
            margin: "0 2px"
          },
          {
            padding: ".35em .625em .75em"
          }
        ]
      },
      legend: {
        padding: "0",
        border: "0",
        display: "block",
        width: "100%",
        fallbacks: [
          {
            border: "0"
          },
          {
            padding: "0"
          }
        ],
        marginBottom: 20,
        fontSize: 21,
        lineHeight: "inherit",
        color: "#333",
        borderBottom: "1px solid #e5e5e5"
      },
      textarea: {
        overflow: "auto"
      },
      optgroup: {
        fontWeight: "bold"
      },
      table: {
        borderSpacing: "0",
        borderCollapse: "collapse",
        backgroundColor: "transparent"
      },
      "td, th": {
        padding: "0"
      },
      "@media print": {
        "*, *:before, *:after": {
          color: "#000 !important",
          textShadow: "none !important",
          background: "transparent !important",
          W: "none !important",
          boxShadow: "none !important"
        },
        "a, a:visited": {
          textDecoration: "underline"
        },
        "a[href]:after": {
          content: '" (" attr(href) ")"'
        },
        "abbr[title]:after": {
          content: '" (" attr(title) ")"'
        },
        'a[href^="#"]:after, a[href^="javascript:"]:after': {
          content: '""'
        },
        "pre, blockquote": {
          border: "1px solid #999",
          pageBreakInside: "avoid"
        },
        thead: {
          display: "table-header-group"
        },
        "tr, img": {
          pageBreakInside: "avoid"
        },
        img: {
          maxWidth: "100% !important"
        },
        "p, h2, h3": {
          orphans: "3",
          widows: "3"
        },
        "h2, h3": {
          pageBreakAfter: "avoid"
        },
        ".navbar": {
          display: "none"
        },
        ".btn > .caret, .dropup > .btn > .caret": {
          borderTopColor: "#000 !important"
        },
        ".label": {
          border: "1px solid #000"
        },
        ".table": {
          borderCollapse: "collapse !important"
        },
        ".table td, .table th": {
          backgroundColor: "#fff !important"
        },
        ".table-bordered th, .table-bordered td": {
          border: "1px solid #ddd !important"
        },
        ".visible-print": {
          display: "block !important"
        },
        "table.visible-print": {
          display: "table !important"
        },
        "tr.visible-print": {
          display: "table-row !important"
        },
        "th.visible-print, td.visible-print": {
          display: "table-cell !important"
        },
        ".visible-print-block": {
          display: "block !important"
        },
        ".visible-print-inline": {
          display: "inline !important"
        },
        ".visible-print-inline-block": {
          display: "inline-block !important"
        },
        ".hidden-print": {
          display: "none !important"
        }
      },
      "@font-face": {
        fontFamily: "'Glyphicons Halflings'",
        src:
          "url('../fonts/glyphicons-halflings-regular.eot?#iefix') format('embedded-opentype'), url('../fonts/glyphicons-halflings-regular.woff2') format('woff2'), url('../fonts/glyphicons-halflings-regular.woff') format('woff'), url('../fonts/glyphicons-halflings-regular.ttf') format('truetype'), url('../fonts/glyphicons-halflings-regular.svg#glyphicons_halflingsregular') format('svg')"
      },
      ".glyphicon": {
        position: "relative",
        top: 1,
        display: "inline-block",
        fontFamily: "'Glyphicons Halflings'",
        fontStyle: "normal",
        fontWeight: "normal",
        lineHeight: "1",
        W: "antialiased",
        M: "grayscale"
      },
      ".glyphicon-asterisk:before": {
        content: '"\\002a"'
      },
      ".glyphicon-plus:before": {
        content: '"\\002b"'
      },
      ".glyphicon-euro:before, .glyphicon-eur:before": {
        content: '"\\20ac"'
      },
      ".glyphicon-minus:before": {
        content: '"\\2212"'
      },
      ".glyphicon-cloud:before": {
        content: '"\\2601"'
      },
      ".glyphicon-envelope:before": {
        content: '"\\2709"'
      },
      ".glyphicon-pencil:before": {
        content: '"\\270f"'
      },
      ".glyphicon-glass:before": {
        content: '"\\e001"'
      },
      ".glyphicon-music:before": {
        content: '"\\e002"'
      },
      ".glyphicon-search:before": {
        content: '"\\e003"'
      },
      ".glyphicon-heart:before": {
        content: '"\\e005"'
      },
      ".glyphicon-star:before": {
        content: '"\\e006"'
      },
      ".glyphicon-star-empty:before": {
        content: '"\\e007"'
      },
      ".glyphicon-user:before": {
        content: '"\\e008"'
      },
      ".glyphicon-film:before": {
        content: '"\\e009"'
      },
      ".glyphicon-th-large:before": {
        content: '"\\e010"'
      },
      ".glyphicon-th:before": {
        content: '"\\e011"'
      },
      ".glyphicon-th-list:before": {
        content: '"\\e012"'
      },
      ".glyphicon-ok:before": {
        content: '"\\e013"'
      },
      ".glyphicon-remove:before": {
        content: '"\\e014"'
      },
      ".glyphicon-zoom-in:before": {
        content: '"\\e015"'
      },
      ".glyphicon-zoom-out:before": {
        content: '"\\e016"'
      },
      ".glyphicon-off:before": {
        content: '"\\e017"'
      },
      ".glyphicon-signal:before": {
        content: '"\\e018"'
      },
      ".glyphicon-cog:before": {
        content: '"\\e019"'
      },
      ".glyphicon-trash:before": {
        content: '"\\e020"'
      },
      ".glyphicon-home:before": {
        content: '"\\e021"'
      },
      ".glyphicon-file:before": {
        content: '"\\e022"'
      },
      ".glyphicon-time:before": {
        content: '"\\e023"'
      },
      ".glyphicon-road:before": {
        content: '"\\e024"'
      },
      ".glyphicon-download-alt:before": {
        content: '"\\e025"'
      },
      ".glyphicon-download:before": {
        content: '"\\e026"'
      },
      ".glyphicon-upload:before": {
        content: '"\\e027"'
      },
      ".glyphicon-inbox:before": {
        content: '"\\e028"'
      },
      ".glyphicon-play-circle:before": {
        content: '"\\e029"'
      },
      ".glyphicon-repeat:before": {
        content: '"\\e030"'
      },
      ".glyphicon-refresh:before": {
        content: '"\\e031"'
      },
      ".glyphicon-list-alt:before": {
        content: '"\\e032"'
      },
      ".glyphicon-lock:before": {
        content: '"\\e033"'
      },
      ".glyphicon-flag:before": {
        content: '"\\e034"'
      },
      ".glyphicon-headphones:before": {
        content: '"\\e035"'
      },
      ".glyphicon-volume-off:before": {
        content: '"\\e036"'
      },
      ".glyphicon-volume-down:before": {
        content: '"\\e037"'
      },
      ".glyphicon-volume-up:before": {
        content: '"\\e038"'
      },
      ".glyphicon-qrcode:before": {
        content: '"\\e039"'
      },
      ".glyphicon-barcode:before": {
        content: '"\\e040"'
      },
      ".glyphicon-tag:before": {
        content: '"\\e041"'
      },
      ".glyphicon-tags:before": {
        content: '"\\e042"'
      },
      ".glyphicon-book:before": {
        content: '"\\e043"'
      },
      ".glyphicon-bookmark:before": {
        content: '"\\e044"'
      },
      ".glyphicon-print:before": {
        content: '"\\e045"'
      },
      ".glyphicon-camera:before": {
        content: '"\\e046"'
      },
      ".glyphicon-font:before": {
        content: '"\\e047"'
      },
      ".glyphicon-bold:before": {
        content: '"\\e048"'
      },
      ".glyphicon-italic:before": {
        content: '"\\e049"'
      },
      ".glyphicon-text-height:before": {
        content: '"\\e050"'
      },
      ".glyphicon-text-width:before": {
        content: '"\\e051"'
      },
      ".glyphicon-align-left:before": {
        content: '"\\e052"'
      },
      ".glyphicon-align-center:before": {
        content: '"\\e053"'
      },
      ".glyphicon-align-right:before": {
        content: '"\\e054"'
      },
      ".glyphicon-align-justify:before": {
        content: '"\\e055"'
      },
      ".glyphicon-list:before": {
        content: '"\\e056"'
      },
      ".glyphicon-indent-left:before": {
        content: '"\\e057"'
      },
      ".glyphicon-indent-right:before": {
        content: '"\\e058"'
      },
      ".glyphicon-facetime-video:before": {
        content: '"\\e059"'
      },
      ".glyphicon-picture:before": {
        content: '"\\e060"'
      },
      ".glyphicon-map-marker:before": {
        content: '"\\e062"'
      },
      ".glyphicon-adjust:before": {
        content: '"\\e063"'
      },
      ".glyphicon-tint:before": {
        content: '"\\e064"'
      },
      ".glyphicon-edit:before": {
        content: '"\\e065"'
      },
      ".glyphicon-share:before": {
        content: '"\\e066"'
      },
      ".glyphicon-check:before": {
        content: '"\\e067"'
      },
      ".glyphicon-move:before": {
        content: '"\\e068"'
      },
      ".glyphicon-step-backward:before": {
        content: '"\\e069"'
      },
      ".glyphicon-fast-backward:before": {
        content: '"\\e070"'
      },
      ".glyphicon-backward:before": {
        content: '"\\e071"'
      },
      ".glyphicon-play:before": {
        content: '"\\e072"'
      },
      ".glyphicon-pause:before": {
        content: '"\\e073"'
      },
      ".glyphicon-stop:before": {
        content: '"\\e074"'
      },
      ".glyphicon-forward:before": {
        content: '"\\e075"'
      },
      ".glyphicon-fast-forward:before": {
        content: '"\\e076"'
      },
      ".glyphicon-step-forward:before": {
        content: '"\\e077"'
      },
      ".glyphicon-eject:before": {
        content: '"\\e078"'
      },
      ".glyphicon-chevron-left:before": {
        content: '"\\e079"'
      },
      ".glyphicon-chevron-right:before": {
        content: '"\\e080"'
      },
      ".glyphicon-plus-sign:before": {
        content: '"\\e081"'
      },
      ".glyphicon-minus-sign:before": {
        content: '"\\e082"'
      },
      ".glyphicon-remove-sign:before": {
        content: '"\\e083"'
      },
      ".glyphicon-ok-sign:before": {
        content: '"\\e084"'
      },
      ".glyphicon-question-sign:before": {
        content: '"\\e085"'
      },
      ".glyphicon-info-sign:before": {
        content: '"\\e086"'
      },
      ".glyphicon-screenshot:before": {
        content: '"\\e087"'
      },
      ".glyphicon-remove-circle:before": {
        content: '"\\e088"'
      },
      ".glyphicon-ok-circle:before": {
        content: '"\\e089"'
      },
      ".glyphicon-ban-circle:before": {
        content: '"\\e090"'
      },
      ".glyphicon-arrow-left:before": {
        content: '"\\e091"'
      },
      ".glyphicon-arrow-right:before": {
        content: '"\\e092"'
      },
      ".glyphicon-arrow-up:before": {
        content: '"\\e093"'
      },
      ".glyphicon-arrow-down:before": {
        content: '"\\e094"'
      },
      ".glyphicon-share-alt:before": {
        content: '"\\e095"'
      },
      ".glyphicon-resize-full:before": {
        content: '"\\e096"'
      },
      ".glyphicon-resize-small:before": {
        content: '"\\e097"'
      },
      ".glyphicon-exclamation-sign:before": {
        content: '"\\e101"'
      },
      ".glyphicon-gift:before": {
        content: '"\\e102"'
      },
      ".glyphicon-leaf:before": {
        content: '"\\e103"'
      },
      ".glyphicon-fire:before": {
        content: '"\\e104"'
      },
      ".glyphicon-eye-open:before": {
        content: '"\\e105"'
      },
      ".glyphicon-eye-close:before": {
        content: '"\\e106"'
      },
      ".glyphicon-warning-sign:before": {
        content: '"\\e107"'
      },
      ".glyphicon-plane:before": {
        content: '"\\e108"'
      },
      ".glyphicon-calendar:before": {
        content: '"\\e109"'
      },
      ".glyphicon-random:before": {
        content: '"\\e110"'
      },
      ".glyphicon-comment:before": {
        content: '"\\e111"'
      },
      ".glyphicon-magnet:before": {
        content: '"\\e112"'
      },
      ".glyphicon-chevron-up:before": {
        content: '"\\e113"'
      },
      ".glyphicon-chevron-down:before": {
        content: '"\\e114"'
      },
      ".glyphicon-retweet:before": {
        content: '"\\e115"'
      },
      ".glyphicon-shopping-cart:before": {
        content: '"\\e116"'
      },
      ".glyphicon-folder-close:before": {
        content: '"\\e117"'
      },
      ".glyphicon-folder-open:before": {
        content: '"\\e118"'
      },
      ".glyphicon-resize-vertical:before": {
        content: '"\\e119"'
      },
      ".glyphicon-resize-horizontal:before": {
        content: '"\\e120"'
      },
      ".glyphicon-hdd:before": {
        content: '"\\e121"'
      },
      ".glyphicon-bullhorn:before": {
        content: '"\\e122"'
      },
      ".glyphicon-bell:before": {
        content: '"\\e123"'
      },
      ".glyphicon-certificate:before": {
        content: '"\\e124"'
      },
      ".glyphicon-thumbs-up:before": {
        content: '"\\e125"'
      },
      ".glyphicon-thumbs-down:before": {
        content: '"\\e126"'
      },
      ".glyphicon-hand-right:before": {
        content: '"\\e127"'
      },
      ".glyphicon-hand-left:before": {
        content: '"\\e128"'
      },
      ".glyphicon-hand-up:before": {
        content: '"\\e129"'
      },
      ".glyphicon-hand-down:before": {
        content: '"\\e130"'
      },
      ".glyphicon-circle-arrow-right:before": {
        content: '"\\e131"'
      },
      ".glyphicon-circle-arrow-left:before": {
        content: '"\\e132"'
      },
      ".glyphicon-circle-arrow-up:before": {
        content: '"\\e133"'
      },
      ".glyphicon-circle-arrow-down:before": {
        content: '"\\e134"'
      },
      ".glyphicon-globe:before": {
        content: '"\\e135"'
      },
      ".glyphicon-wrench:before": {
        content: '"\\e136"'
      },
      ".glyphicon-tasks:before": {
        content: '"\\e137"'
      },
      ".glyphicon-filter:before": {
        content: '"\\e138"'
      },
      ".glyphicon-briefcase:before": {
        content: '"\\e139"'
      },
      ".glyphicon-fullscreen:before": {
        content: '"\\e140"'
      },
      ".glyphicon-dashboard:before": {
        content: '"\\e141"'
      },
      ".glyphicon-paperclip:before": {
        content: '"\\e142"'
      },
      ".glyphicon-heart-empty:before": {
        content: '"\\e143"'
      },
      ".glyphicon-link:before": {
        content: '"\\e144"'
      },
      ".glyphicon-phone:before": {
        content: '"\\e145"'
      },
      ".glyphicon-pushpin:before": {
        content: '"\\e146"'
      },
      ".glyphicon-usd:before": {
        content: '"\\e148"'
      },
      ".glyphicon-gbp:before": {
        content: '"\\e149"'
      },
      ".glyphicon-sort:before": {
        content: '"\\e150"'
      },
      ".glyphicon-sort-by-alphabet:before": {
        content: '"\\e151"'
      },
      ".glyphicon-sort-by-alphabet-alt:before": {
        content: '"\\e152"'
      },
      ".glyphicon-sort-by-order:before": {
        content: '"\\e153"'
      },
      ".glyphicon-sort-by-order-alt:before": {
        content: '"\\e154"'
      },
      ".glyphicon-sort-by-attributes:before": {
        content: '"\\e155"'
      },
      ".glyphicon-sort-by-attributes-alt:before": {
        content: '"\\e156"'
      },
      ".glyphicon-unchecked:before": {
        content: '"\\e157"'
      },
      ".glyphicon-expand:before": {
        content: '"\\e158"'
      },
      ".glyphicon-collapse-down:before": {
        content: '"\\e159"'
      },
      ".glyphicon-collapse-up:before": {
        content: '"\\e160"'
      },
      ".glyphicon-log-in:before": {
        content: '"\\e161"'
      },
      ".glyphicon-flash:before": {
        content: '"\\e162"'
      },
      ".glyphicon-log-out:before": {
        content: '"\\e163"'
      },
      ".glyphicon-new-window:before": {
        content: '"\\e164"'
      },
      ".glyphicon-record:before": {
        content: '"\\e165"'
      },
      ".glyphicon-save:before": {
        content: '"\\e166"'
      },
      ".glyphicon-open:before": {
        content: '"\\e167"'
      },
      ".glyphicon-saved:before": {
        content: '"\\e168"'
      },
      ".glyphicon-import:before": {
        content: '"\\e169"'
      },
      ".glyphicon-export:before": {
        content: '"\\e170"'
      },
      ".glyphicon-send:before": {
        content: '"\\e171"'
      },
      ".glyphicon-floppy-disk:before": {
        content: '"\\e172"'
      },
      ".glyphicon-floppy-saved:before": {
        content: '"\\e173"'
      },
      ".glyphicon-floppy-remove:before": {
        content: '"\\e174"'
      },
      ".glyphicon-floppy-save:before": {
        content: '"\\e175"'
      },
      ".glyphicon-floppy-open:before": {
        content: '"\\e176"'
      },
      ".glyphicon-credit-card:before": {
        content: '"\\e177"'
      },
      ".glyphicon-transfer:before": {
        content: '"\\e178"'
      },
      ".glyphicon-cutlery:before": {
        content: '"\\e179"'
      },
      ".glyphicon-header:before": {
        content: '"\\e180"'
      },
      ".glyphicon-compressed:before": {
        content: '"\\e181"'
      },
      ".glyphicon-earphone:before": {
        content: '"\\e182"'
      },
      ".glyphicon-phone-alt:before": {
        content: '"\\e183"'
      },
      ".glyphicon-tower:before": {
        content: '"\\e184"'
      },
      ".glyphicon-stats:before": {
        content: '"\\e185"'
      },
      ".glyphicon-sd-video:before": {
        content: '"\\e186"'
      },
      ".glyphicon-hd-video:before": {
        content: '"\\e187"'
      },
      ".glyphicon-subtitles:before": {
        content: '"\\e188"'
      },
      ".glyphicon-sound-stereo:before": {
        content: '"\\e189"'
      },
      ".glyphicon-sound-dolby:before": {
        content: '"\\e190"'
      },
      ".glyphicon-sound-5-1:before": {
        content: '"\\e191"'
      },
      ".glyphicon-sound-6-1:before": {
        content: '"\\e192"'
      },
      ".glyphicon-sound-7-1:before": {
        content: '"\\e193"'
      },
      ".glyphicon-copyright-mark:before": {
        content: '"\\e194"'
      },
      ".glyphicon-registration-mark:before": {
        content: '"\\e195"'
      },
      ".glyphicon-cloud-download:before": {
        content: '"\\e197"'
      },
      ".glyphicon-cloud-upload:before": {
        content: '"\\e198"'
      },
      ".glyphicon-tree-conifer:before": {
        content: '"\\e199"'
      },
      ".glyphicon-tree-deciduous:before": {
        content: '"\\e200"'
      },
      ".glyphicon-cd:before": {
        content: '"\\e201"'
      },
      ".glyphicon-save-file:before": {
        content: '"\\e202"'
      },
      ".glyphicon-open-file:before": {
        content: '"\\e203"'
      },
      ".glyphicon-level-up:before": {
        content: '"\\e204"'
      },
      ".glyphicon-copy:before": {
        content: '"\\e205"'
      },
      ".glyphicon-paste:before": {
        content: '"\\e206"'
      },
      ".glyphicon-alert:before": {
        content: '"\\e209"'
      },
      ".glyphicon-equalizer:before": {
        content: '"\\e210"'
      },
      ".glyphicon-king:before": {
        content: '"\\e211"'
      },
      ".glyphicon-queen:before": {
        content: '"\\e212"'
      },
      ".glyphicon-pawn:before": {
        content: '"\\e213"'
      },
      ".glyphicon-bishop:before": {
        content: '"\\e214"'
      },
      ".glyphicon-knight:before": {
        content: '"\\e215"'
      },
      ".glyphicon-baby-formula:before": {
        content: '"\\e216"'
      },
      ".glyphicon-tent:before": {
        content: '"\\26fa"'
      },
      ".glyphicon-blackboard:before": {
        content: '"\\e218"'
      },
      ".glyphicon-bed:before": {
        content: '"\\e219"'
      },
      ".glyphicon-apple:before": {
        content: '"\\f8ff"'
      },
      ".glyphicon-erase:before": {
        content: '"\\e221"'
      },
      ".glyphicon-hourglass:before": {
        content: '"\\231b"'
      },
      ".glyphicon-lamp:before": {
        content: '"\\e223"'
      },
      ".glyphicon-duplicate:before": {
        content: '"\\e224"'
      },
      ".glyphicon-piggy-bank:before": {
        content: '"\\e225"'
      },
      ".glyphicon-scissors:before": {
        content: '"\\e226"'
      },
      ".glyphicon-bitcoin:before": {
        content: '"\\e227"'
      },
      ".glyphicon-btc:before": {
        content: '"\\e227"'
      },
      ".glyphicon-xbt:before": {
        content: '"\\e227"'
      },
      ".glyphicon-yen:before": {
        content: '"\\00a5"'
      },
      ".glyphicon-jpy:before": {
        content: '"\\00a5"'
      },
      ".glyphicon-ruble:before": {
        content: '"\\20bd"'
      },
      ".glyphicon-rub:before": {
        content: '"\\20bd"'
      },
      ".glyphicon-scale:before": {
        content: '"\\e230"'
      },
      ".glyphicon-ice-lolly:before": {
        content: '"\\e231"'
      },
      ".glyphicon-ice-lolly-tasted:before": {
        content: '"\\e232"'
      },
      ".glyphicon-education:before": {
        content: '"\\e233"'
      },
      ".glyphicon-option-horizontal:before": {
        content: '"\\e234"'
      },
      ".glyphicon-option-vertical:before": {
        content: '"\\e235"'
      },
      ".glyphicon-menu-hamburger:before": {
        content: '"\\e236"'
      },
      ".glyphicon-modal-window:before": {
        content: '"\\e237"'
      },
      ".glyphicon-oil:before": {
        content: '"\\e238"'
      },
      ".glyphicon-grain:before": {
        content: '"\\e239"'
      },
      ".glyphicon-sunglasses:before": {
        content: '"\\e240"'
      },
      ".glyphicon-text-size:before": {
        content: '"\\e241"'
      },
      ".glyphicon-text-color:before": {
        content: '"\\e242"'
      },
      ".glyphicon-text-background:before": {
        content: '"\\e243"'
      },
      ".glyphicon-object-align-top:before": {
        content: '"\\e244"'
      },
      ".glyphicon-object-align-bottom:before": {
        content: '"\\e245"'
      },
      ".glyphicon-object-align-horizontal:before": {
        content: '"\\e246"'
      },
      ".glyphicon-object-align-left:before": {
        content: '"\\e247"'
      },
      ".glyphicon-object-align-vertical:before": {
        content: '"\\e248"'
      },
      ".glyphicon-object-align-right:before": {
        content: '"\\e249"'
      },
      ".glyphicon-triangle-right:before": {
        content: '"\\e250"'
      },
      ".glyphicon-triangle-left:before": {
        content: '"\\e251"'
      },
      ".glyphicon-triangle-bottom:before": {
        content: '"\\e252"'
      },
      ".glyphicon-triangle-top:before": {
        content: '"\\e253"'
      },
      ".glyphicon-console:before": {
        content: '"\\e254"'
      },
      ".glyphicon-superscript:before": {
        content: '"\\e255"'
      },
      ".glyphicon-subscript:before": {
        content: '"\\e256"'
      },
      ".glyphicon-menu-left:before": {
        content: '"\\e257"'
      },
      ".glyphicon-menu-right:before": {
        content: '"\\e258"'
      },
      ".glyphicon-menu-down:before": {
        content: '"\\e259"'
      },
      ".glyphicon-menu-up:before": {
        content: '"\\e260"'
      },
      "*": {
        W: "border-box",
        M: "border-box",
        boxSizing: "border-box"
      },
      "*:before, *:after": {
        W: "border-box",
        M: "border-box",
        boxSizing: "border-box"
      },
      "input, button, select, textarea": {
        fontFamily: "inherit",
        fontSize: "inherit",
        lineHeight: "inherit"
      },
      "a:hover, a:focus": {
        color: "#23527c",
        textDecoration: "underline"
      },
      "a:focus": {
        outline: "5px auto -webkit-focus-ring-color",
        outlineOffset: -2
      },
      ".img-responsive, .thumbnail > img, .thumbnail a > img, .carousel-inner > .item > img, .carousel-inner > .item > a > img": {
        display: "block",
        maxWidth: "100%",
        height: "auto"
      },
      ".img-rounded": {
        borderRadius: 6
      },
      ".img-thumbnail": {
        display: "inline-block",
        maxWidth: "100%",
        height: "auto",
        padding: 4,
        lineHeight: "1.42857143",
        backgroundColor: "#fff",
        border: "1px solid #ddd",
        borderRadius: 4,
        W: "all .2s ease-in-out",
        O: "all .2s ease-in-out",
        transition: "all .2s ease-in-out"
      },
      ".img-circle": {
        borderRadius: "50%"
      },
      ".sr-only": {
        position: "absolute",
        width: 1,
        height: 1,
        padding: "0",
        margin: -1,
        overflow: "hidden",
        clip: "rect(0, 0, 0, 0)",
        border: "0"
      },
      ".sr-only-focusable:active, .sr-only-focusable:focus": {
        position: "static",
        width: "auto",
        height: "auto",
        margin: "0",
        overflow: "visible",
        clip: "auto"
      },
      '[role="button"]': {
        cursor: "pointer"
      },
      "h1, h2, h3, h4, h5, h6, .h1, .h2, .h3, .h4, .h5, .h6": {
        fontFamily: "inherit",
        fontWeight: "500",
        lineHeight: "1.1",
        color: "inherit"
      },
      "h1 small, h2 small, h3 small, h4 small, h5 small, h6 small, .h1 small, .h2 small, .h3 small, .h4 small, .h5 small, .h6 small, h1 .small, h2 .small, h3 .small, h4 .small, h5 .small, h6 .small, .h1 .small, .h2 .small, .h3 .small, .h4 .small, .h5 .small, .h6 .small": {
        fontWeight: "normal",
        lineHeight: "1",
        color: "#777"
      },
      "h1, .h1, h2, .h2, h3, .h3": {
        marginTop: 20,
        marginBottom: 10
      },
      "h1 small, .h1 small, h2 small, .h2 small, h3 small, .h3 small, h1 .small, .h1 .small, h2 .small, .h2 .small, h3 .small, .h3 .small": {
        fontSize: "65%"
      },
      "h4, .h4, h5, .h5, h6, .h6": {
        marginTop: 10,
        marginBottom: 10
      },
      "h4 small, .h4 small, h5 small, .h5 small, h6 small, .h6 small, h4 .small, .h4 .small, h5 .small, .h5 .small, h6 .small, .h6 .small": {
        fontSize: "75%"
      },
      "h1, .h1": {
        fontSize: 36
      },
      "h2, .h2": {
        fontSize: 30
      },
      "h3, .h3": {
        fontSize: 24
      },
      "h4, .h4": {
        fontSize: 18
      },
      "h5, .h5": {
        fontSize: 14
      },
      "h6, .h6": {
        fontSize: 12
      },
      p: {
        margin: "0 0 10px"
      },
      ".lead": {
        marginBottom: 20,
        fontSize: 16,
        fontWeight: "300",
        lineHeight: "1.4"
      },
      "@media (min-width: 768px)": {
        ".lead": {
          fontSize: 21
        },
        ".dl-horizontal dt": {
          float: "left",
          width: 160,
          overflow: "hidden",
          clear: "left",
          textAlign: "right",
          textOverflow: "ellipsis",
          whiteSpace: "nowrap"
        },
        ".dl-horizontal dd": {
          marginLeft: 180
        },
        ".container": {
          width: 750
        },
        ".col-sm-1, .col-sm-2, .col-sm-3, .col-sm-4, .col-sm-5, .col-sm-6, .col-sm-7, .col-sm-8, .col-sm-9, .col-sm-10, .col-sm-11, .col-sm-12": {
          float: "left"
        },
        ".col-sm-12": {
          width: "100%"
        },
        ".col-sm-11": {
          width: "91.66666667%"
        },
        ".col-sm-10": {
          width: "83.33333333%"
        },
        ".col-sm-9": {
          width: "75%"
        },
        ".col-sm-8": {
          width: "66.66666667%"
        },
        ".col-sm-7": {
          width: "58.33333333%"
        },
        ".col-sm-6": {
          width: "50%"
        },
        ".col-sm-5": {
          width: "41.66666667%"
        },
        ".col-sm-4": {
          width: "33.33333333%"
        },
        ".col-sm-3": {
          width: "25%"
        },
        ".col-sm-2": {
          width: "16.66666667%"
        },
        ".col-sm-1": {
          width: "8.33333333%"
        },
        ".col-sm-pull-12": {
          right: "100%"
        },
        ".col-sm-pull-11": {
          right: "91.66666667%"
        },
        ".col-sm-pull-10": {
          right: "83.33333333%"
        },
        ".col-sm-pull-9": {
          right: "75%"
        },
        ".col-sm-pull-8": {
          right: "66.66666667%"
        },
        ".col-sm-pull-7": {
          right: "58.33333333%"
        },
        ".col-sm-pull-6": {
          right: "50%"
        },
        ".col-sm-pull-5": {
          right: "41.66666667%"
        },
        ".col-sm-pull-4": {
          right: "33.33333333%"
        },
        ".col-sm-pull-3": {
          right: "25%"
        },
        ".col-sm-pull-2": {
          right: "16.66666667%"
        },
        ".col-sm-pull-1": {
          right: "8.33333333%"
        },
        ".col-sm-pull-0": {
          right: "auto"
        },
        ".col-sm-push-12": {
          left: "100%"
        },
        ".col-sm-push-11": {
          left: "91.66666667%"
        },
        ".col-sm-push-10": {
          left: "83.33333333%"
        },
        ".col-sm-push-9": {
          left: "75%"
        },
        ".col-sm-push-8": {
          left: "66.66666667%"
        },
        ".col-sm-push-7": {
          left: "58.33333333%"
        },
        ".col-sm-push-6": {
          left: "50%"
        },
        ".col-sm-push-5": {
          left: "41.66666667%"
        },
        ".col-sm-push-4": {
          left: "33.33333333%"
        },
        ".col-sm-push-3": {
          left: "25%"
        },
        ".col-sm-push-2": {
          left: "16.66666667%"
        },
        ".col-sm-push-1": {
          left: "8.33333333%"
        },
        ".col-sm-push-0": {
          left: "auto"
        },
        ".col-sm-offset-12": {
          marginLeft: "100%"
        },
        ".col-sm-offset-11": {
          marginLeft: "91.66666667%"
        },
        ".col-sm-offset-10": {
          marginLeft: "83.33333333%"
        },
        ".col-sm-offset-9": {
          marginLeft: "75%"
        },
        ".col-sm-offset-8": {
          marginLeft: "66.66666667%"
        },
        ".col-sm-offset-7": {
          marginLeft: "58.33333333%"
        },
        ".col-sm-offset-6": {
          marginLeft: "50%"
        },
        ".col-sm-offset-5": {
          marginLeft: "41.66666667%"
        },
        ".col-sm-offset-4": {
          marginLeft: "33.33333333%"
        },
        ".col-sm-offset-3": {
          marginLeft: "25%"
        },
        ".col-sm-offset-2": {
          marginLeft: "16.66666667%"
        },
        ".col-sm-offset-1": {
          marginLeft: "8.33333333%"
        },
        ".col-sm-offset-0": {
          marginLeft: "0"
        },
        ".form-inline .form-group": {
          display: "inline-block",
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".form-inline .form-control": {
          display: "inline-block",
          width: "auto",
          verticalAlign: "middle"
        },
        ".form-inline .form-control-static": {
          display: "inline-block"
        },
        ".form-inline .input-group": {
          display: "inline-table",
          verticalAlign: "middle"
        },
        ".form-inline .input-group .input-group-addon, .form-inline .input-group .input-group-btn, .form-inline .input-group .form-control": {
          width: "auto"
        },
        ".form-inline .input-group > .form-control": {
          width: "100%"
        },
        ".form-inline .control-label": {
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".form-inline .radio, .form-inline .checkbox": {
          display: "inline-block",
          marginTop: "0",
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".form-inline .radio label, .form-inline .checkbox label": {
          paddingLeft: "0"
        },
        '.form-inline .radio input[type="radio"], .form-inline .checkbox input[type="checkbox"]': {
          position: "relative",
          marginLeft: "0"
        },
        ".form-inline .has-feedback .form-control-feedback": {
          top: "0"
        },
        ".form-horizontal .control-label": {
          paddingTop: 7,
          marginBottom: "0",
          textAlign: "right"
        },
        ".form-horizontal .form-group-lg .control-label": {
          paddingTop: 11,
          fontSize: 18
        },
        ".form-horizontal .form-group-sm .control-label": {
          paddingTop: 6,
          fontSize: 12
        },
        ".navbar-right .dropdown-menu": {
          right: "0",
          left: "auto"
        },
        ".navbar-right .dropdown-menu-left": {
          right: "auto",
          left: "0"
        },
        ".nav-tabs.nav-justified > li": {
          display: "table-cell",
          width: "1%"
        },
        ".nav-tabs.nav-justified > li > a": {
          marginBottom: "0",
          borderBottom: "1px solid #ddd",
          borderRadius: "4px 4px 0 0"
        },
        ".nav-tabs.nav-justified > .active > a, .nav-tabs.nav-justified > .active > a:hover, .nav-tabs.nav-justified > .active > a:focus": {
          borderBottomColor: "#fff"
        },
        ".nav-justified > li": {
          display: "table-cell",
          width: "1%"
        },
        ".nav-justified > li > a": {
          marginBottom: "0"
        },
        ".nav-tabs-justified > li > a": {
          borderBottom: "1px solid #ddd",
          borderRadius: "4px 4px 0 0"
        },
        ".nav-tabs-justified > .active > a, .nav-tabs-justified > .active > a:hover, .nav-tabs-justified > .active > a:focus": {
          borderBottomColor: "#fff"
        },
        ".navbar": {
          borderRadius: 4
        },
        ".navbar-header": {
          float: "left"
        },
        ".navbar-collapse": {
          width: "auto",
          borderTop: "0",
          W: "none",
          boxShadow: "none"
        },
        ".navbar-collapse.collapse": {
          display: "block !important",
          height: "auto !important",
          paddingBottom: "0",
          overflow: "visible !important"
        },
        ".navbar-collapse.in": {
          overflowY: "visible"
        },
        ".navbar-fixed-top .navbar-collapse, .navbar-static-top .navbar-collapse, .navbar-fixed-bottom .navbar-collapse": {
          paddingRight: "0",
          paddingLeft: "0"
        },
        ".container > .navbar-header, .container-fluid > .navbar-header, .container > .navbar-collapse, .container-fluid > .navbar-collapse": {
          marginRight: "0",
          marginLeft: "0"
        },
        ".navbar-static-top": {
          borderRadius: "0"
        },
        ".navbar-fixed-top, .navbar-fixed-bottom": {
          borderRadius: "0"
        },
        ".navbar > .container .navbar-brand, .navbar > .container-fluid .navbar-brand": {
          marginLeft: -15
        },
        ".navbar-toggle": {
          display: "none"
        },
        ".navbar-nav": {
          float: "left",
          margin: "0"
        },
        ".navbar-nav > li": {
          float: "left"
        },
        ".navbar-nav > li > a": {
          paddingTop: 15,
          paddingBottom: 15
        },
        ".navbar-form .form-group": {
          display: "inline-block",
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".navbar-form .form-control": {
          display: "inline-block",
          width: "auto",
          verticalAlign: "middle"
        },
        ".navbar-form .form-control-static": {
          display: "inline-block"
        },
        ".navbar-form .input-group": {
          display: "inline-table",
          verticalAlign: "middle"
        },
        ".navbar-form .input-group .input-group-addon, .navbar-form .input-group .input-group-btn, .navbar-form .input-group .form-control": {
          width: "auto"
        },
        ".navbar-form .input-group > .form-control": {
          width: "100%"
        },
        ".navbar-form .control-label": {
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".navbar-form .radio, .navbar-form .checkbox": {
          display: "inline-block",
          marginTop: "0",
          marginBottom: "0",
          verticalAlign: "middle"
        },
        ".navbar-form .radio label, .navbar-form .checkbox label": {
          paddingLeft: "0"
        },
        '.navbar-form .radio input[type="radio"], .navbar-form .checkbox input[type="checkbox"]': {
          position: "relative",
          marginLeft: "0"
        },
        ".navbar-form .has-feedback .form-control-feedback": {
          top: "0"
        },
        ".navbar-form": {
          width: "auto",
          paddingTop: "0",
          paddingBottom: "0",
          marginRight: "0",
          marginLeft: "0",
          border: "0",
          W: "none",
          boxShadow: "none"
        },
        ".navbar-text": {
          float: "left",
          marginRight: 15,
          marginLeft: 15
        },
        ".navbar-left": {
          float: "left !important"
        },
        ".navbar-right": {
          float: "right !important",
          marginRight: -15
        },
        ".navbar-right ~ .navbar-right": {
          marginRight: "0"
        },
        ".modal-dialog": {
          width: 600,
          margin: "30px auto"
        },
        ".modal-content": {
          W: "0 5px 15px rgba(0, 0, 0, .5)",
          boxShadow: "0 5px 15px rgba(0, 0, 0, .5)"
        },
        ".modal-sm": {
          width: 300
        }
      },
      "small, .small": {
        fontSize: "85%"
      },
      "mark, .mark": {
        padding: ".2em",
        backgroundColor: "#fcf8e3"
      },
      ".text-left": {
        textAlign: "left"
      },
      ".text-right": {
        textAlign: "right"
      },
      ".text-center": {
        textAlign: "center"
      },
      ".text-justify": {
        textAlign: "justify"
      },
      ".text-nowrap": {
        whiteSpace: "nowrap"
      },
      ".text-lowercase": {
        textTransform: "lowercase"
      },
      ".text-uppercase": {
        textTransform: "uppercase"
      },
      ".text-capitalize": {
        textTransform: "capitalize"
      },
      ".text-muted": {
        color: "#777"
      },
      ".text-primary": {
        color: theme.colors.primary.main
      },
      "a.text-primary:hover, a.text-primary:focus": {
        color: theme.colors.primary.dark
      },
      ".text-success": {
        color: "#3c763d"
      },
      "a.text-success:hover, a.text-success:focus": {
        color: "#2b542c"
      },
      ".text-info": {
        color: "#31708f"
      },
      "a.text-info:hover, a.text-info:focus": {
        color: "#245269"
      },
      ".text-warning": {
        color: "#8a6d3b"
      },
      "a.text-warning:hover, a.text-warning:focus": {
        color: "#66512c"
      },
      ".text-danger": {
        color: "#a94442"
      },
      "a.text-danger:hover, a.text-danger:focus": {
        color: "#843534"
      },
      ".bg-primary": {
        color: "#fff",
        backgroundColor: theme.colors.primary.main
      },
      "a.bg-primary:hover, a.bg-primary:focus": {
        backgroundColor: theme.colors.primary.dark
      },
      ".bg-success": {
        backgroundColor: "#dff0d8"
      },
      "a.bg-success:hover, a.bg-success:focus": {
        backgroundColor: "#c1e2b3"
      },
      ".bg-info": {
        backgroundColor: "#d9edf7"
      },
      "a.bg-info:hover, a.bg-info:focus": {
        backgroundColor: "#afd9ee"
      },
      ".bg-warning": {
        backgroundColor: "#fcf8e3"
      },
      "a.bg-warning:hover, a.bg-warning:focus": {
        backgroundColor: "#f7ecb5"
      },
      ".bg-danger": {
        backgroundColor: "#f2dede"
      },
      "a.bg-danger:hover, a.bg-danger:focus": {
        backgroundColor: "#e4b9b9"
      },
      ".page-header": {
        paddingBottom: 9,
        margin: "40px 0 20px",
        borderBottom: "1px solid #eee"
      },
      "ul, ol": {
        marginTop: "0",
        marginBottom: 10
      },
      "ul ul, ol ul, ul ol, ol ol": {
        marginBottom: "0"
      },
      ".list-unstyled": {
        paddingLeft: "0",
        listStyle: "none"
      },
      ".list-inline": {
        paddingLeft: "0",
        marginLeft: -5,
        listStyle: "none"
      },
      ".list-inline > li": {
        display: "inline-block",
        paddingRight: 5,
        paddingLeft: 5
      },
      dl: {
        marginTop: "0",
        marginBottom: 20
      },
      "dt, dd": {
        lineHeight: "1.42857143"
      },
      dt: {
        fontWeight: "bold"
      },
      dd: {
        marginLeft: "0"
      },
      "abbr[title], abbr[data-original-title]": {
        cursor: "help",
        borderBottom: "1px dotted #777"
      },
      ".initialism": {
        fontSize: "90%",
        textTransform: "uppercase"
      },
      blockquote: {
        padding: "10px 20px",
        margin: "0 0 20px",
        fontSize: 17.5,
        borderLeft: "5px solid #eee"
      },
      "blockquote p:last-child, blockquote ul:last-child, blockquote ol:last-child": {
        marginBottom: "0"
      },
      "blockquote footer, blockquote small, blockquote .small": {
        display: "block",
        fontSize: "80%",
        lineHeight: "1.42857143",
        color: "#777"
      },
      "blockquote footer:before, blockquote small:before, blockquote .small:before": {
        content: "'\\2014 \\00A0'"
      },
      ".blockquote-reverse, blockquote.pull-right": {
        paddingRight: 15,
        paddingLeft: "0",
        textAlign: "right",
        borderRight: "5px solid #eee",
        borderLeft: "0"
      },
      ".blockquote-reverse footer:before, blockquote.pull-right footer:before, .blockquote-reverse small:before, blockquote.pull-right small:before, .blockquote-reverse .small:before, blockquote.pull-right .small:before": {
        content: "''"
      },
      ".blockquote-reverse footer:after, blockquote.pull-right footer:after, .blockquote-reverse small:after, blockquote.pull-right small:after, .blockquote-reverse .small:after, blockquote.pull-right .small:after": {
        content: "'\\00A0 \\2014'"
      },
      address: {
        marginBottom: 20,
        fontStyle: "normal",
        lineHeight: "1.42857143"
      },
      code: {
        padding: "2px 4px",
        fontSize: "90%",
        color: "#c7254e",
        backgroundColor: "#f9f2f4",
        borderRadius: 4
      },
      kbd: {
        padding: "2px 4px",
        fontSize: "90%",
        color: "#fff",
        backgroundColor: "#333",
        borderRadius: 3,
        W: "inset 0 -1px 0 rgba(0, 0, 0, .25)",
        boxShadow: "inset 0 -1px 0 rgba(0, 0, 0, .25)"
      },
      "kbd kbd": {
        padding: "0",
        fontSize: "100%",
        fontWeight: "bold",
        W: "none",
        boxShadow: "none"
      },
      "pre code": {
        padding: "0",
        fontSize: "inherit",
        color: "inherit",
        whiteSpace: "pre-wrap",
        backgroundColor: "transparent",
        borderRadius: "0"
      },
      ".pre-scrollable": {
        maxHeight: 340,
        overflowY: "scroll"
      },
      ".container": {
        paddingRight: 15,
        paddingLeft: 15,
        marginRight: "auto",
        marginLeft: "auto"
      },
      "@media (min-width: 992px)": {
        ".container": {
          width: 970
        },
        ".col-md-1, .col-md-2, .col-md-3, .col-md-4, .col-md-5, .col-md-6, .col-md-7, .col-md-8, .col-md-9, .col-md-10, .col-md-11, .col-md-12": {
          float: "left"
        },
        ".col-md-12": {
          width: "100%"
        },
        ".col-md-11": {
          width: "91.66666667%"
        },
        ".col-md-10": {
          width: "83.33333333%"
        },
        ".col-md-9": {
          width: "75%"
        },
        ".col-md-8": {
          width: "66.66666667%"
        },
        ".col-md-7": {
          width: "58.33333333%"
        },
        ".col-md-6": {
          width: "50%"
        },
        ".col-md-5": {
          width: "41.66666667%"
        },
        ".col-md-4": {
          width: "33.33333333%"
        },
        ".col-md-3": {
          width: "25%"
        },
        ".col-md-2": {
          width: "16.66666667%"
        },
        ".col-md-1": {
          width: "8.33333333%"
        },
        ".col-md-pull-12": {
          right: "100%"
        },
        ".col-md-pull-11": {
          right: "91.66666667%"
        },
        ".col-md-pull-10": {
          right: "83.33333333%"
        },
        ".col-md-pull-9": {
          right: "75%"
        },
        ".col-md-pull-8": {
          right: "66.66666667%"
        },
        ".col-md-pull-7": {
          right: "58.33333333%"
        },
        ".col-md-pull-6": {
          right: "50%"
        },
        ".col-md-pull-5": {
          right: "41.66666667%"
        },
        ".col-md-pull-4": {
          right: "33.33333333%"
        },
        ".col-md-pull-3": {
          right: "25%"
        },
        ".col-md-pull-2": {
          right: "16.66666667%"
        },
        ".col-md-pull-1": {
          right: "8.33333333%"
        },
        ".col-md-pull-0": {
          right: "auto"
        },
        ".col-md-push-12": {
          left: "100%"
        },
        ".col-md-push-11": {
          left: "91.66666667%"
        },
        ".col-md-push-10": {
          left: "83.33333333%"
        },
        ".col-md-push-9": {
          left: "75%"
        },
        ".col-md-push-8": {
          left: "66.66666667%"
        },
        ".col-md-push-7": {
          left: "58.33333333%"
        },
        ".col-md-push-6": {
          left: "50%"
        },
        ".col-md-push-5": {
          left: "41.66666667%"
        },
        ".col-md-push-4": {
          left: "33.33333333%"
        },
        ".col-md-push-3": {
          left: "25%"
        },
        ".col-md-push-2": {
          left: "16.66666667%"
        },
        ".col-md-push-1": {
          left: "8.33333333%"
        },
        ".col-md-push-0": {
          left: "auto"
        },
        ".col-md-offset-12": {
          marginLeft: "100%"
        },
        ".col-md-offset-11": {
          marginLeft: "91.66666667%"
        },
        ".col-md-offset-10": {
          marginLeft: "83.33333333%"
        },
        ".col-md-offset-9": {
          marginLeft: "75%"
        },
        ".col-md-offset-8": {
          marginLeft: "66.66666667%"
        },
        ".col-md-offset-7": {
          marginLeft: "58.33333333%"
        },
        ".col-md-offset-6": {
          marginLeft: "50%"
        },
        ".col-md-offset-5": {
          marginLeft: "41.66666667%"
        },
        ".col-md-offset-4": {
          marginLeft: "33.33333333%"
        },
        ".col-md-offset-3": {
          marginLeft: "25%"
        },
        ".col-md-offset-2": {
          marginLeft: "16.66666667%"
        },
        ".col-md-offset-1": {
          marginLeft: "8.33333333%"
        },
        ".col-md-offset-0": {
          marginLeft: "0"
        },
        ".modal-lg": {
          width: 900
        }
      },
      [theme.breakpoints.up("lg")]: {
        ".container": {
          width: theme.breakpoints.width("lg") - 30
        },
        ".col-lg-1, .col-lg-2, .col-lg-3, .col-lg-4, .col-lg-5, .col-lg-6, .col-lg-7, .col-lg-8, .col-lg-9, .col-lg-10, .col-lg-11, .col-lg-12": {
          float: "left"
        },
        ".col-lg-12": {
          width: "100%"
        },
        ".col-lg-11": {
          width: "91.66666667%"
        },
        ".col-lg-10": {
          width: "83.33333333%"
        },
        ".col-lg-9": {
          width: "75%"
        },
        ".col-lg-8": {
          width: "66.66666667%"
        },
        ".col-lg-7": {
          width: "58.33333333%"
        },
        ".col-lg-6": {
          width: "50%"
        },
        ".col-lg-5": {
          width: "41.66666667%"
        },
        ".col-lg-4": {
          width: "33.33333333%"
        },
        ".col-lg-3": {
          width: "25%"
        },
        ".col-lg-2": {
          width: "16.66666667%"
        },
        ".col-lg-1": {
          width: "8.33333333%"
        },
        ".col-lg-pull-12": {
          right: "100%"
        },
        ".col-lg-pull-11": {
          right: "91.66666667%"
        },
        ".col-lg-pull-10": {
          right: "83.33333333%"
        },
        ".col-lg-pull-9": {
          right: "75%"
        },
        ".col-lg-pull-8": {
          right: "66.66666667%"
        },
        ".col-lg-pull-7": {
          right: "58.33333333%"
        },
        ".col-lg-pull-6": {
          right: "50%"
        },
        ".col-lg-pull-5": {
          right: "41.66666667%"
        },
        ".col-lg-pull-4": {
          right: "33.33333333%"
        },
        ".col-lg-pull-3": {
          right: "25%"
        },
        ".col-lg-pull-2": {
          right: "16.66666667%"
        },
        ".col-lg-pull-1": {
          right: "8.33333333%"
        },
        ".col-lg-pull-0": {
          right: "auto"
        },
        ".col-lg-push-12": {
          left: "100%"
        },
        ".col-lg-push-11": {
          left: "91.66666667%"
        },
        ".col-lg-push-10": {
          left: "83.33333333%"
        },
        ".col-lg-push-9": {
          left: "75%"
        },
        ".col-lg-push-8": {
          left: "66.66666667%"
        },
        ".col-lg-push-7": {
          left: "58.33333333%"
        },
        ".col-lg-push-6": {
          left: "50%"
        },
        ".col-lg-push-5": {
          left: "41.66666667%"
        },
        ".col-lg-push-4": {
          left: "33.33333333%"
        },
        ".col-lg-push-3": {
          left: "25%"
        },
        ".col-lg-push-2": {
          left: "16.66666667%"
        },
        ".col-lg-push-1": {
          left: "8.33333333%"
        },
        ".col-lg-push-0": {
          left: "auto"
        },
        ".col-lg-offset-12": {
          marginLeft: "100%"
        },
        ".col-lg-offset-11": {
          marginLeft: "91.66666667%"
        },
        ".col-lg-offset-10": {
          marginLeft: "83.33333333%"
        },
        ".col-lg-offset-9": {
          marginLeft: "75%"
        },
        ".col-lg-offset-8": {
          marginLeft: "66.66666667%"
        },
        ".col-lg-offset-7": {
          marginLeft: "58.33333333%"
        },
        ".col-lg-offset-6": {
          marginLeft: "50%"
        },
        ".col-lg-offset-5": {
          marginLeft: "41.66666667%"
        },
        ".col-lg-offset-4": {
          marginLeft: "33.33333333%"
        },
        ".col-lg-offset-3": {
          marginLeft: "25%"
        },
        ".col-lg-offset-2": {
          marginLeft: "16.66666667%"
        },
        ".col-lg-offset-1": {
          marginLeft: "8.33333333%"
        },
        ".col-lg-offset-0": {
          marginLeft: "0"
        },
        ".visible-lg": {
          display: "block !important"
        },
        "table.visible-lg": {
          display: "table !important"
        },
        "tr.visible-lg": {
          display: "table-row !important"
        },
        "th.visible-lg, td.visible-lg": {
          display: "table-cell !important"
        },
        ".visible-lg-block": {
          display: "block !important"
        },
        ".visible-lg-inline": {
          display: "inline !important"
        },
        ".visible-lg-inline-block": {
          display: "inline-block !important"
        },
        ".hidden-lg": {
          display: "none !important"
        }
      },
      ".container-fluid": {
        paddingRight: 15,
        paddingLeft: 15,
        marginRight: "auto",
        marginLeft: "auto"
      },
      ".row": {
        marginRight: -15,
        marginLeft: -15
      },
      ".col-xs-1, .col-sm-1, .col-md-1, .col-lg-1, .col-xs-2, .col-sm-2, .col-md-2, .col-lg-2, .col-xs-3, .col-sm-3, .col-md-3, .col-lg-3, .col-xs-4, .col-sm-4, .col-md-4, .col-lg-4, .col-xs-5, .col-sm-5, .col-md-5, .col-lg-5, .col-xs-6, .col-sm-6, .col-md-6, .col-lg-6, .col-xs-7, .col-sm-7, .col-md-7, .col-lg-7, .col-xs-8, .col-sm-8, .col-md-8, .col-lg-8, .col-xs-9, .col-sm-9, .col-md-9, .col-lg-9, .col-xs-10, .col-sm-10, .col-md-10, .col-lg-10, .col-xs-11, .col-sm-11, .col-md-11, .col-lg-11, .col-xs-12, .col-sm-12, .col-md-12, .col-lg-12": {
        position: "relative",
        minHeight: 1,
        paddingRight: 15,
        paddingLeft: 15
      },
      ".col-xs-1, .col-xs-2, .col-xs-3, .col-xs-4, .col-xs-5, .col-xs-6, .col-xs-7, .col-xs-8, .col-xs-9, .col-xs-10, .col-xs-11, .col-xs-12": {
        float: "left"
      },
      ".col-xs-12": {
        width: "100%"
      },
      ".col-xs-11": {
        width: "91.66666667%"
      },
      ".col-xs-10": {
        width: "83.33333333%"
      },
      ".col-xs-9": {
        width: "75%"
      },
      ".col-xs-8": {
        width: "66.66666667%"
      },
      ".col-xs-7": {
        width: "58.33333333%"
      },
      ".col-xs-6": {
        width: "50%"
      },
      ".col-xs-5": {
        width: "41.66666667%"
      },
      ".col-xs-4": {
        width: "33.33333333%"
      },
      ".col-xs-3": {
        width: "25%"
      },
      ".col-xs-2": {
        width: "16.66666667%"
      },
      ".col-xs-1": {
        width: "8.33333333%"
      },
      ".col-xs-pull-12": {
        right: "100%"
      },
      ".col-xs-pull-11": {
        right: "91.66666667%"
      },
      ".col-xs-pull-10": {
        right: "83.33333333%"
      },
      ".col-xs-pull-9": {
        right: "75%"
      },
      ".col-xs-pull-8": {
        right: "66.66666667%"
      },
      ".col-xs-pull-7": {
        right: "58.33333333%"
      },
      ".col-xs-pull-6": {
        right: "50%"
      },
      ".col-xs-pull-5": {
        right: "41.66666667%"
      },
      ".col-xs-pull-4": {
        right: "33.33333333%"
      },
      ".col-xs-pull-3": {
        right: "25%"
      },
      ".col-xs-pull-2": {
        right: "16.66666667%"
      },
      ".col-xs-pull-1": {
        right: "8.33333333%"
      },
      ".col-xs-pull-0": {
        right: "auto"
      },
      ".col-xs-push-12": {
        left: "100%"
      },
      ".col-xs-push-11": {
        left: "91.66666667%"
      },
      ".col-xs-push-10": {
        left: "83.33333333%"
      },
      ".col-xs-push-9": {
        left: "75%"
      },
      ".col-xs-push-8": {
        left: "66.66666667%"
      },
      ".col-xs-push-7": {
        left: "58.33333333%"
      },
      ".col-xs-push-6": {
        left: "50%"
      },
      ".col-xs-push-5": {
        left: "41.66666667%"
      },
      ".col-xs-push-4": {
        left: "33.33333333%"
      },
      ".col-xs-push-3": {
        left: "25%"
      },
      ".col-xs-push-2": {
        left: "16.66666667%"
      },
      ".col-xs-push-1": {
        left: "8.33333333%"
      },
      ".col-xs-push-0": {
        left: "auto"
      },
      ".col-xs-offset-12": {
        marginLeft: "100%"
      },
      ".col-xs-offset-11": {
        marginLeft: "91.66666667%"
      },
      ".col-xs-offset-10": {
        marginLeft: "83.33333333%"
      },
      ".col-xs-offset-9": {
        marginLeft: "75%"
      },
      ".col-xs-offset-8": {
        marginLeft: "66.66666667%"
      },
      ".col-xs-offset-7": {
        marginLeft: "58.33333333%"
      },
      ".col-xs-offset-6": {
        marginLeft: "50%"
      },
      ".col-xs-offset-5": {
        marginLeft: "41.66666667%"
      },
      ".col-xs-offset-4": {
        marginLeft: "33.33333333%"
      },
      ".col-xs-offset-3": {
        marginLeft: "25%"
      },
      ".col-xs-offset-2": {
        marginLeft: "16.66666667%"
      },
      ".col-xs-offset-1": {
        marginLeft: "8.33333333%"
      },
      ".col-xs-offset-0": {
        marginLeft: "0"
      },
      caption: {
        paddingTop: 8,
        paddingBottom: 8,
        color: "#777",
        textAlign: "left"
      },
      th: {
        textAlign: "left"
      },
      ".table": {
        width: "100%",
        maxWidth: "100%",
        marginBottom: 20
      },
      ".table > thead > tr > th, .table > tbody > tr > th, .table > tfoot > tr > th, .table > thead > tr > td, .table > tbody > tr > td, .table > tfoot > tr > td": {
        padding: 8,
        lineHeight: "1.42857143",
        verticalAlign: "top",
        borderTop: "1px solid #ddd"
      },
      ".table > thead > tr > th": {
        verticalAlign: "bottom",
        borderBottom: "2px solid #ddd"
      },
      ".table > caption + thead > tr:first-child > th, .table > colgroup + thead > tr:first-child > th, .table > thead:first-child > tr:first-child > th, .table > caption + thead > tr:first-child > td, .table > colgroup + thead > tr:first-child > td, .table > thead:first-child > tr:first-child > td": {
        borderTop: "0"
      },
      ".table > tbody + tbody": {
        borderTop: "2px solid #ddd"
      },
      ".table .table": {
        backgroundColor: "#fff"
      },
      ".table-condensed > thead > tr > th, .table-condensed > tbody > tr > th, .table-condensed > tfoot > tr > th, .table-condensed > thead > tr > td, .table-condensed > tbody > tr > td, .table-condensed > tfoot > tr > td": {
        padding: 5
      },
      ".table-bordered": {
        border: "1px solid #ddd"
      },
      ".table-bordered > thead > tr > th, .table-bordered > tbody > tr > th, .table-bordered > tfoot > tr > th, .table-bordered > thead > tr > td, .table-bordered > tbody > tr > td, .table-bordered > tfoot > tr > td": {
        border: "1px solid #ddd"
      },
      ".table-bordered > thead > tr > th, .table-bordered > thead > tr > td": {
        borderBottomWidth: 2
      },
      ".table-striped > tbody > tr:nth-of-type(odd)": {
        backgroundColor: "#f9f9f9"
      },
      ".table-hover > tbody > tr:hover": {
        backgroundColor: "#f5f5f5"
      },
      'table col[class*="col-"]': {
        position: "static",
        display: "table-column",
        float: "none"
      },
      'table td[class*="col-"], table th[class*="col-"]': {
        position: "static",
        display: "table-cell",
        float: "none"
      },
      ".table > thead > tr > td.active, .table > tbody > tr > td.active, .table > tfoot > tr > td.active, .table > thead > tr > th.active, .table > tbody > tr > th.active, .table > tfoot > tr > th.active, .table > thead > tr.active > td, .table > tbody > tr.active > td, .table > tfoot > tr.active > td, .table > thead > tr.active > th, .table > tbody > tr.active > th, .table > tfoot > tr.active > th": {
        backgroundColor: "#f5f5f5"
      },
      ".table-hover > tbody > tr > td.active:hover, .table-hover > tbody > tr > th.active:hover, .table-hover > tbody > tr.active:hover > td, .table-hover > tbody > tr:hover > .active, .table-hover > tbody > tr.active:hover > th": {
        backgroundColor: "#e8e8e8"
      },
      ".table > thead > tr > td.success, .table > tbody > tr > td.success, .table > tfoot > tr > td.success, .table > thead > tr > th.success, .table > tbody > tr > th.success, .table > tfoot > tr > th.success, .table > thead > tr.success > td, .table > tbody > tr.success > td, .table > tfoot > tr.success > td, .table > thead > tr.success > th, .table > tbody > tr.success > th, .table > tfoot > tr.success > th": {
        backgroundColor: "#dff0d8"
      },
      ".table-hover > tbody > tr > td.success:hover, .table-hover > tbody > tr > th.success:hover, .table-hover > tbody > tr.success:hover > td, .table-hover > tbody > tr:hover > .success, .table-hover > tbody > tr.success:hover > th": {
        backgroundColor: "#d0e9c6"
      },
      ".table > thead > tr > td.info, .table > tbody > tr > td.info, .table > tfoot > tr > td.info, .table > thead > tr > th.info, .table > tbody > tr > th.info, .table > tfoot > tr > th.info, .table > thead > tr.info > td, .table > tbody > tr.info > td, .table > tfoot > tr.info > td, .table > thead > tr.info > th, .table > tbody > tr.info > th, .table > tfoot > tr.info > th": {
        backgroundColor: "#d9edf7"
      },
      ".table-hover > tbody > tr > td.info:hover, .table-hover > tbody > tr > th.info:hover, .table-hover > tbody > tr.info:hover > td, .table-hover > tbody > tr:hover > .info, .table-hover > tbody > tr.info:hover > th": {
        backgroundColor: "#c4e3f3"
      },
      ".table > thead > tr > td.warning, .table > tbody > tr > td.warning, .table > tfoot > tr > td.warning, .table > thead > tr > th.warning, .table > tbody > tr > th.warning, .table > tfoot > tr > th.warning, .table > thead > tr.warning > td, .table > tbody > tr.warning > td, .table > tfoot > tr.warning > td, .table > thead > tr.warning > th, .table > tbody > tr.warning > th, .table > tfoot > tr.warning > th": {
        backgroundColor: "#fcf8e3"
      },
      ".table-hover > tbody > tr > td.warning:hover, .table-hover > tbody > tr > th.warning:hover, .table-hover > tbody > tr.warning:hover > td, .table-hover > tbody > tr:hover > .warning, .table-hover > tbody > tr.warning:hover > th": {
        backgroundColor: "#faf2cc"
      },
      ".table > thead > tr > td.danger, .table > tbody > tr > td.danger, .table > tfoot > tr > td.danger, .table > thead > tr > th.danger, .table > tbody > tr > th.danger, .table > tfoot > tr > th.danger, .table > thead > tr.danger > td, .table > tbody > tr.danger > td, .table > tfoot > tr.danger > td, .table > thead > tr.danger > th, .table > tbody > tr.danger > th, .table > tfoot > tr.danger > th": {
        backgroundColor: "#f2dede"
      },
      ".table-hover > tbody > tr > td.danger:hover, .table-hover > tbody > tr > th.danger:hover, .table-hover > tbody > tr.danger:hover > td, .table-hover > tbody > tr:hover > .danger, .table-hover > tbody > tr.danger:hover > th": {
        backgroundColor: "#ebcccc"
      },
      ".table-responsive": {
        minHeight: ".01%",
        overflowX: "auto"
      },
      "@media screen and (max-width: 767px)": {
        ".table-responsive": {
          width: "100%",
          marginBottom: 15,
          overflowY: "hidden",
          M: "-ms-autohiding-scrollbar",
          border: "1px solid #ddd"
        },
        ".table-responsive > .table": {
          marginBottom: "0"
        },
        ".table-responsive > .table > thead > tr > th, .table-responsive > .table > tbody > tr > th, .table-responsive > .table > tfoot > tr > th, .table-responsive > .table > thead > tr > td, .table-responsive > .table > tbody > tr > td, .table-responsive > .table > tfoot > tr > td": {
          whiteSpace: "nowrap"
        },
        ".table-responsive > .table-bordered": {
          border: "0"
        },
        ".table-responsive > .table-bordered > thead > tr > th:first-child, .table-responsive > .table-bordered > tbody > tr > th:first-child, .table-responsive > .table-bordered > tfoot > tr > th:first-child, .table-responsive > .table-bordered > thead > tr > td:first-child, .table-responsive > .table-bordered > tbody > tr > td:first-child, .table-responsive > .table-bordered > tfoot > tr > td:first-child": {
          borderLeft: "0"
        },
        ".table-responsive > .table-bordered > thead > tr > th:last-child, .table-responsive > .table-bordered > tbody > tr > th:last-child, .table-responsive > .table-bordered > tfoot > tr > th:last-child, .table-responsive > .table-bordered > thead > tr > td:last-child, .table-responsive > .table-bordered > tbody > tr > td:last-child, .table-responsive > .table-bordered > tfoot > tr > td:last-child": {
          borderRight: "0"
        },
        ".table-responsive > .table-bordered > tbody > tr:last-child > th, .table-responsive > .table-bordered > tfoot > tr:last-child > th, .table-responsive > .table-bordered > tbody > tr:last-child > td, .table-responsive > .table-bordered > tfoot > tr:last-child > td": {
          borderBottom: "0"
        }
      },
      label: {
        ...theme.typography.caption,
        fontWeight: 600,
        display: "inline-block",
        maxWidth: "100%",
        marginBottom: theme.spacing / 2
      },
      'input[type="radio"], input[type="checkbox"]': {
        margin: "4px 0 0",
        lineHeight: "normal"
      },
      'input[type="file"]': {
        display: "block"
      },
      'input[type="range"]': {
        display: "block",
        width: "100%"
      },
      "select[multiple], select[size]": {
        height: "auto"
      },
      'input[type="file"]:focus, input[type="radio"]:focus, input[type="checkbox"]:focus': {
        outline: "5px auto -webkit-focus-ring-color",
        outlineOffset: -2
      },
      output: {
        display: "block",
        paddingTop: 7,
        fontSize: 14,
        lineHeight: "1.42857143",
        color: "#555"
      },
      ".form-control": {
        ...theme.typography.body,
        outline: 0,
        display: "block",
        width: "100%",
        height: theme.spacing * 4,
        padding: 0,
        backgroundColor: "#fff",
        backgroundImage: "none",
        border: "none",
        borderRadius: 0,
        transitionDuration: theme.transition.time,
        position: "relative" as "relative",
        borderBottom: `1px solid ${theme.colors.disabled}`,
        boxShadow: `0 0 0 transparent`,
        "&:focus": {
          borderBottomColor: theme.colors.secondary.main,
          boxShadow: `0 1px ${theme.colors.secondary.main}`
        }
      },
      ".form-control::-moz-placeholder": {
        color: "#999",
        opacity: "1"
      },
      ".form-control:-ms-input-placeholder": {
        color: "#999"
      },
      ".form-control::-webkit-input-placeholder": {
        color: "#999"
      },
      ".form-control::-ms-expand": {
        backgroundColor: "transparent",
        border: "0"
      },
      ".form-control[disabled], .form-control[readonly], fieldset[disabled] .form-control": {
        backgroundColor: "#eee",
        opacity: "1"
      },
      ".form-control[disabled], fieldset[disabled] .form-control": {
        cursor: "not-allowed"
      },
      "textarea.form-control": {
        height: "auto",
        minHeight: "10rem",
        resize: "vertical"
      },
      "@media screen and (-webkit-min-device-pixel-ratio: 0)": {
        'input[type="date"].form-control, input[type="time"].form-control, input[type="datetime-local"].form-control, input[type="month"].form-control': {
          lineHeight: 34
        },
        'input[type="date"].input-sm, input[type="time"].input-sm, input[type="datetime-local"].input-sm, input[type="month"].input-sm, .input-group-sm input[type="date"], .input-group-sm input[type="time"], .input-group-sm input[type="datetime-local"], .input-group-sm input[type="month"]': {
          lineHeight: 30
        },
        'input[type="date"].input-lg, input[type="time"].input-lg, input[type="datetime-local"].input-lg, input[type="month"].input-lg, .input-group-lg input[type="date"], .input-group-lg input[type="time"], .input-group-lg input[type="datetime-local"], .input-group-lg input[type="month"]': {
          lineHeight: 46
        }
      },
      ".form-group": {
        marginBottom: 15
      },
      ".radio, .checkbox": {
        position: "relative",
        display: "block",
        marginTop: 10,
        marginBottom: 10
      },
      ".radio label, .checkbox label": {
        minHeight: 20,
        paddingLeft: 20,
        marginBottom: "0",
        fontWeight: "normal",
        cursor: "pointer"
      },
      ".checkbox": {
        "& label": {
          padding: '10px 30px'
        },
        '& input[type="checkbox"]': {
          "&:checked": {
            backgroundColor: theme.colors.secondary.main
          },
          "&:focus, &:active": {
            boxShadow: `0 0 2px 1px ${theme.colors.secondary.dark}`
          },
          border: `2px ${theme.colors.secondary.main} solid`,
          borderRadius: 5,
          cursor: "pointer" as "pointer",
          height: 20,
          lineHeight: 20,
          margin: "0 10px 0 -30px",
          outline: "none",
          position: "absolute" as "absolute",
          top: 10,
          transition: theme.transition.time,
          "-webkit-appearance": "none",
          width: 20
        }
      },
      ".radio + .radio, .checkbox + .checkbox": {
        marginTop: -5
      },
      ".radio-inline, .checkbox-inline": {
        position: "relative",
        display: "inline-block",
        paddingLeft: 20,
        marginBottom: "0",
        fontWeight: "normal",
        verticalAlign: "middle",
        cursor: "pointer"
      },
      ".radio-inline + .radio-inline, .checkbox-inline + .checkbox-inline": {
        marginTop: "0",
        marginLeft: 10
      },
      'input[type="radio"][disabled], input[type="checkbox"][disabled], input[type="radio"].disabled, input[type="checkbox"].disabled, fieldset[disabled] input[type="radio"], fieldset[disabled] input[type="checkbox"]': {
        cursor: "not-allowed"
      },
      ".radio-inline.disabled, .checkbox-inline.disabled, fieldset[disabled] .radio-inline, fieldset[disabled] .checkbox-inline": {
        cursor: "not-allowed"
      },
      ".radio.disabled label, .checkbox.disabled label, fieldset[disabled] .radio label, fieldset[disabled] .checkbox label": {
        cursor: "not-allowed"
      },
      ".form-control-static": {
        minHeight: 34,
        paddingTop: 7,
        paddingBottom: 7,
        marginBottom: "0"
      },
      ".form-control-static.input-lg, .form-control-static.input-sm": {
        paddingRight: "0",
        paddingLeft: "0"
      },
      ".input-sm": {
        height: 30,
        padding: "5px 10px",
        fontSize: 12,
        lineHeight: "1.5",
        borderRadius: 3
      },
      "select.input-sm": {
        height: 30,
        lineHeight: 30
      },
      "textarea.input-sm, select[multiple].input-sm": {
        height: "auto"
      },
      ".form-group-sm .form-control": {
        height: 30,
        padding: "5px 10px",
        fontSize: 12,
        lineHeight: "1.5",
        borderRadius: 3
      },
      ".form-group-sm select.form-control": {
        height: 30,
        lineHeight: 30
      },
      ".form-group-sm textarea.form-control, .form-group-sm select[multiple].form-control": {
        height: "auto"
      },
      ".form-group-sm .form-control-static": {
        height: 30,
        minHeight: 32,
        padding: "6px 10px",
        fontSize: 12,
        lineHeight: "1.5"
      },
      ".input-lg": {
        height: 46,
        padding: "10px 16px",
        fontSize: 18,
        lineHeight: "1.3333333",
        borderRadius: 6
      },
      "select.input-lg": {
        height: 46,
        lineHeight: 46
      },
      "textarea.input-lg, select[multiple].input-lg": {
        height: "auto"
      },
      ".form-group-lg .form-control": {
        height: 46,
        padding: "10px 16px",
        fontSize: 18,
        lineHeight: "1.3333333",
        borderRadius: 6
      },
      ".form-group-lg select.form-control": {
        height: 46,
        lineHeight: 46
      },
      ".form-group-lg textarea.form-control, .form-group-lg select[multiple].form-control": {
        height: "auto"
      },
      ".form-group-lg .form-control-static": {
        height: 46,
        minHeight: 38,
        padding: "11px 16px",
        fontSize: 18,
        lineHeight: "1.3333333"
      },
      ".has-feedback": {
        position: "relative"
      },
      ".has-feedback .form-control": {
        paddingRight: 42.5
      },
      ".form-control-feedback": {
        position: "absolute",
        top: "0",
        right: "0",
        zIndex: "2",
        display: "block",
        width: 34,
        height: 34,
        lineHeight: 34,
        textAlign: "center",
        pointerEvents: "none"
      },
      ".input-lg + .form-control-feedback, .input-group-lg + .form-control-feedback, .form-group-lg .form-control + .form-control-feedback": {
        width: 46,
        height: 46,
        lineHeight: 46
      },
      ".input-sm + .form-control-feedback, .input-group-sm + .form-control-feedback, .form-group-sm .form-control + .form-control-feedback": {
        width: 30,
        height: 30,
        lineHeight: 30
      },
      ".has-success .help-block, .has-success .control-label, .has-success .radio, .has-success .checkbox, .has-success .radio-inline, .has-success .checkbox-inline, .has-success.radio label, .has-success.checkbox label, .has-success.radio-inline label, .has-success.checkbox-inline label": {
        color: "#3c763d"
      },
      ".has-success .form-control": {
        borderColor: "#3c763d",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075)",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075)"
      },
      ".has-success .form-control:focus": {
        borderColor: "#2b542c",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #67b168",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #67b168"
      },
      ".has-success .input-group-addon": {
        color: "#3c763d",
        backgroundColor: "#dff0d8",
        borderColor: "#3c763d"
      },
      ".has-success .form-control-feedback": {
        color: "#3c763d"
      },
      ".has-warning .help-block, .has-warning .control-label, .has-warning .radio, .has-warning .checkbox, .has-warning .radio-inline, .has-warning .checkbox-inline, .has-warning.radio label, .has-warning.checkbox label, .has-warning.radio-inline label, .has-warning.checkbox-inline label": {
        color: "#8a6d3b"
      },
      ".has-warning .form-control": {
        borderColor: "#8a6d3b",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075)",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075)"
      },
      ".has-warning .form-control:focus": {
        borderColor: "#66512c",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #c0a16b",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #c0a16b"
      },
      ".has-warning .input-group-addon": {
        color: "#8a6d3b",
        backgroundColor: "#fcf8e3",
        borderColor: "#8a6d3b"
      },
      ".has-warning .form-control-feedback": {
        color: "#8a6d3b"
      },
      ".has-error .help-block, .has-error .control-label, .has-error .radio, .has-error .checkbox, .has-error .radio-inline, .has-error .checkbox-inline, .has-error.radio label, .has-error.checkbox label, .has-error.radio-inline label, .has-error.checkbox-inline label": {
        color: "#a94442"
      },
      ".has-error .form-control": {
        borderColor: "#a94442",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075)",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075)"
      },
      ".has-error .form-control:focus": {
        borderColor: "#843534",
        W: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #ce8483",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .075), 0 0 6px #ce8483"
      },
      ".has-error .input-group-addon": {
        color: "#a94442",
        backgroundColor: "#f2dede",
        borderColor: "#a94442"
      },
      ".has-error .form-control-feedback": {
        color: "#a94442"
      },
      ".has-feedback label ~ .form-control-feedback": {
        top: 25
      },
      ".has-feedback label.sr-only ~ .form-control-feedback": {
        top: "0"
      },
      ".help-block": {
        display: "block",
        marginTop: 5,
        marginBottom: 10,
        color: "#737373"
      },
      ".form-horizontal .radio, .form-horizontal .checkbox, .form-horizontal .radio-inline, .form-horizontal .checkbox-inline": {
        paddingTop: 7,
        marginTop: "0",
        marginBottom: "0"
      },
      ".form-horizontal .radio, .form-horizontal .checkbox": {
        minHeight: 27
      },
      ".form-horizontal .form-group": {
        marginRight: -15,
        marginLeft: -15
      },
      ".form-horizontal .has-feedback .form-control-feedback": {
        right: 15
      },
      ".btn": {
        ...theme.typography.button,
        transitionDuration: theme.transition.time,
        display: "inline-block",
        padding: `${theme.spacing}px ${theme.spacing * 2}px`,
        marginBottom: "0",
        textAlign: "center",
        whiteSpace: "nowrap",
        verticalAlign: "middle",
        M: "none",
        touchAction: "manipulation",
        cursor: "pointer",
        W: "none",
        fallbacks: [
          {
            M: "none"
          },
          {
            M: "manipulation"
          }
        ],
        userSelect: "none",
        backgroundImage: "none",
        border: "0px solid transparent",
        borderRadius: 2
      },
      ".btn:focus, .btn:active:focus, .btn.active:focus, .btn.focus, .btn:active.focus, .btn.active.focus": {
        outline: "5px auto -webkit-focus-ring-color",
        outlineOffset: -2
      },
      ".btn:hover, .btn:focus, .btn.focus": {
        color: "#333",
        textDecoration: "none"
      },
      ".btn:active, .btn.active": {
        backgroundImage: "none",
        outline: "0",
        W: "inset 0 3px 5px rgba(0, 0, 0, .125)",
        boxShadow: "inset 0 3px 5px rgba(0, 0, 0, .125)"
      },
      ".btn.disabled, .btn[disabled], fieldset[disabled] .btn": {
        cursor: "not-allowed",
        filter: "alpha(opacity=65)",
        W: "none",
        boxShadow: "none",
        opacity: ".65"
      },
      "a.btn.disabled, fieldset[disabled] a.btn": {
        pointerEvents: "none"
      },
      ".btn-default": {
        color: "#333",
        backgroundColor: "#fff",
        borderColor: "#ccc"
      },
      ".btn-default:focus, .btn-default.focus": {
        color: "#333",
        backgroundColor: "#e6e6e6",
        borderColor: "#8c8c8c"
      },
      ".btn-default:hover": {
        color: "#333",
        backgroundColor: "#e6e6e6",
        borderColor: "#adadad"
      },
      ".btn-default:active, .btn-default.active, .open > .dropdown-toggle.btn-default": {
        color: "#333",
        backgroundColor: "#e6e6e6",
        borderColor: "#adadad",
        backgroundImage: "none"
      },
      ".btn-default:active:hover, .btn-default.active:hover, .open > .dropdown-toggle.btn-default:hover, .btn-default:active:focus, .btn-default.active:focus, .open > .dropdown-toggle.btn-default:focus, .btn-default:active.focus, .btn-default.active.focus, .open > .dropdown-toggle.btn-default.focus": {
        color: "#333",
        backgroundColor: "#d4d4d4",
        borderColor: "#8c8c8c"
      },
      ".btn-default.disabled:hover, .btn-default[disabled]:hover, fieldset[disabled] .btn-default:hover, .btn-default.disabled:focus, .btn-default[disabled]:focus, fieldset[disabled] .btn-default:focus, .btn-default.disabled.focus, .btn-default[disabled].focus, fieldset[disabled] .btn-default.focus": {
        backgroundColor: "#fff",
        borderColor: "#ccc"
      },
      ".btn-default .badge": {
        color: "#fff",
        backgroundColor: "#333"
      },
      ".btn-primary": {
        color: "#fff",
        backgroundColor: theme.colors.primary.main,
        borderColor: "#2e6da4"
      },
      ".btn-primary:focus, .btn-primary.focus": {
        color: "#fff",
        backgroundColor: theme.colors.primary.dark,
        borderColor: "#122b40"
      },
      ".btn-primary:hover": {
        color: "#fff",
        backgroundColor: theme.colors.primary.dark,
        borderColor: "#204d74"
      },
      ".btn-primary:active, .btn-primary.active, .open > .dropdown-toggle.btn-primary": {
        color: "#fff",
        backgroundColor: theme.colors.primary.dark,
        borderColor: "#204d74",
        backgroundImage: "none"
      },
      ".btn-primary:active:hover, .btn-primary.active:hover, .open > .dropdown-toggle.btn-primary:hover, .btn-primary:active:focus, .btn-primary.active:focus, .open > .dropdown-toggle.btn-primary:focus, .btn-primary:active.focus, .btn-primary.active.focus, .open > .dropdown-toggle.btn-primary.focus": {
        color: "#fff",
        backgroundColor: "#204d74",
        borderColor: "#122b40"
      },
      ".btn-primary.disabled:hover, .btn-primary[disabled]:hover, fieldset[disabled] .btn-primary:hover, .btn-primary.disabled:focus, .btn-primary[disabled]:focus, fieldset[disabled] .btn-primary:focus, .btn-primary.disabled.focus, .btn-primary[disabled].focus, fieldset[disabled] .btn-primary.focus": {
        backgroundColor: theme.colors.primary.main,
        borderColor: "#2e6da4"
      },
      ".btn-primary .badge": {
        color: theme.colors.primary.main,
        backgroundColor: "#fff"
      },
      ".btn-success": {
        color: "#fff",
        backgroundColor: theme.colors.success.main,
        borderColor: "#4cae4c"
      },
      ".btn-success:focus, .btn-success.focus": {
        color: "#fff",
        backgroundColor: theme.colors.success.dark,
        borderColor: "#255625"
      },
      ".btn-success:hover": {
        color: "#fff",
        backgroundColor: theme.colors.success.dark,
        borderColor: "#398439"
      },
      ".btn-success:active, .btn-success.active, .open > .dropdown-toggle.btn-success": {
        color: "#fff",
        backgroundColor: theme.colors.success.dark,
        borderColor: "#398439",
        backgroundImage: "none"
      },
      ".btn-success:active:hover, .btn-success.active:hover, .open > .dropdown-toggle.btn-success:hover, .btn-success:active:focus, .btn-success.active:focus, .open > .dropdown-toggle.btn-success:focus, .btn-success:active.focus, .btn-success.active.focus, .open > .dropdown-toggle.btn-success.focus": {
        color: "#fff",
        backgroundColor: "#398439",
        borderColor: "#255625"
      },
      ".btn-success.disabled:hover, .btn-success[disabled]:hover, fieldset[disabled] .btn-success:hover, .btn-success.disabled:focus, .btn-success[disabled]:focus, fieldset[disabled] .btn-success:focus, .btn-success.disabled.focus, .btn-success[disabled].focus, fieldset[disabled] .btn-success.focus": {
        backgroundColor: theme.colors.success.main,
        borderColor: "#4cae4c"
      },
      ".btn-success .badge": {
        color: theme.colors.success.main,
        backgroundColor: "#fff"
      },
      ".btn-info": {
        color: "#fff",
        backgroundColor: "#5bc0de",
        borderColor: "#46b8da"
      },
      ".btn-info:focus, .btn-info.focus": {
        color: "#fff",
        backgroundColor: "#31b0d5",
        borderColor: "#1b6d85"
      },
      ".btn-info:hover": {
        color: "#fff",
        backgroundColor: "#31b0d5",
        borderColor: "#269abc"
      },
      ".btn-info:active, .btn-info.active, .open > .dropdown-toggle.btn-info": {
        color: "#fff",
        backgroundColor: "#31b0d5",
        borderColor: "#269abc",
        backgroundImage: "none"
      },
      ".btn-info:active:hover, .btn-info.active:hover, .open > .dropdown-toggle.btn-info:hover, .btn-info:active:focus, .btn-info.active:focus, .open > .dropdown-toggle.btn-info:focus, .btn-info:active.focus, .btn-info.active.focus, .open > .dropdown-toggle.btn-info.focus": {
        color: "#fff",
        backgroundColor: "#269abc",
        borderColor: "#1b6d85"
      },
      ".btn-info.disabled:hover, .btn-info[disabled]:hover, fieldset[disabled] .btn-info:hover, .btn-info.disabled:focus, .btn-info[disabled]:focus, fieldset[disabled] .btn-info:focus, .btn-info.disabled.focus, .btn-info[disabled].focus, fieldset[disabled] .btn-info.focus": {
        backgroundColor: "#5bc0de",
        borderColor: "#46b8da"
      },
      ".btn-info .badge": {
        color: "#5bc0de",
        backgroundColor: "#fff"
      },
      ".btn-warning": {
        color: "#fff",
        backgroundColor: "#f0ad4e",
        borderColor: "#eea236"
      },
      ".btn-warning:focus, .btn-warning.focus": {
        color: "#fff",
        backgroundColor: "#ec971f",
        borderColor: "#985f0d"
      },
      ".btn-warning:hover": {
        color: "#fff",
        backgroundColor: "#ec971f",
        borderColor: "#d58512"
      },
      ".btn-warning:active, .btn-warning.active, .open > .dropdown-toggle.btn-warning": {
        color: "#fff",
        backgroundColor: "#ec971f",
        borderColor: "#d58512",
        backgroundImage: "none"
      },
      ".btn-warning:active:hover, .btn-warning.active:hover, .open > .dropdown-toggle.btn-warning:hover, .btn-warning:active:focus, .btn-warning.active:focus, .open > .dropdown-toggle.btn-warning:focus, .btn-warning:active.focus, .btn-warning.active.focus, .open > .dropdown-toggle.btn-warning.focus": {
        color: "#fff",
        backgroundColor: "#d58512",
        borderColor: "#985f0d"
      },
      ".btn-warning.disabled:hover, .btn-warning[disabled]:hover, fieldset[disabled] .btn-warning:hover, .btn-warning.disabled:focus, .btn-warning[disabled]:focus, fieldset[disabled] .btn-warning:focus, .btn-warning.disabled.focus, .btn-warning[disabled].focus, fieldset[disabled] .btn-warning.focus": {
        backgroundColor: "#f0ad4e",
        borderColor: "#eea236"
      },
      ".btn-warning .badge": {
        color: "#f0ad4e",
        backgroundColor: "#fff"
      },
      ".btn-danger": {
        color: "#fff",
        backgroundColor: theme.colors.error.main,
        borderColor: "#d43f3a"
      },
      ".btn-danger:focus, .btn-danger.focus": {
        color: "#fff",
        backgroundColor: theme.colors.error.dark,
        borderColor: "#761c19"
      },
      ".btn-danger:hover": {
        color: "#fff",
        backgroundColor: theme.colors.error.dark,
        borderColor: "#ac2925"
      },
      ".btn-danger:active, .btn-danger.active, .open > .dropdown-toggle.btn-danger": {
        color: "#fff",
        backgroundColor: theme.colors.error.dark,
        borderColor: "#ac2925",
        backgroundImage: "none"
      },
      ".btn-danger:active:hover, .btn-danger.active:hover, .open > .dropdown-toggle.btn-danger:hover, .btn-danger:active:focus, .btn-danger.active:focus, .open > .dropdown-toggle.btn-danger:focus, .btn-danger:active.focus, .btn-danger.active.focus, .open > .dropdown-toggle.btn-danger.focus": {
        color: "#fff",
        backgroundColor: "#ac2925",
        borderColor: "#761c19"
      },
      ".btn-danger.disabled:hover, .btn-danger[disabled]:hover, fieldset[disabled] .btn-danger:hover, .btn-danger.disabled:focus, .btn-danger[disabled]:focus, fieldset[disabled] .btn-danger:focus, .btn-danger.disabled.focus, .btn-danger[disabled].focus, fieldset[disabled] .btn-danger.focus": {
        backgroundColor: theme.colors.error.main,
        borderColor: "#d43f3a"
      },
      ".btn-danger .badge": {
        color: theme.colors.error.main,
        backgroundColor: "#fff"
      },
      ".btn-link": {
        fontWeight: "normal",
        color: theme.colors.primary.main,
        borderRadius: "0"
      },
      ".btn-link, .btn-link:active, .btn-link.active, .btn-link[disabled], fieldset[disabled] .btn-link": {
        backgroundColor: "transparent",
        W: "none",
        boxShadow: "none"
      },
      ".btn-link, .btn-link:hover, .btn-link:focus, .btn-link:active": {
        borderColor: "transparent"
      },
      ".btn-link:hover, .btn-link:focus": {
        color: "#23527c",
        textDecoration: "underline",
        backgroundColor: "transparent"
      },
      ".btn-link[disabled]:hover, fieldset[disabled] .btn-link:hover, .btn-link[disabled]:focus, fieldset[disabled] .btn-link:focus": {
        color: "#777",
        textDecoration: "none"
      },
      ".btn-lg, .btn-group-lg > .btn": {
        padding: "10px 16px",
        fontSize: 18,
        lineHeight: "1.3333333",
        borderRadius: 6
      },
      ".btn-sm, .btn-group-sm > .btn": {
        padding: "5px 10px",
        fontSize: 12,
        lineHeight: "1.5",
        borderRadius: 3
      },
      ".btn-xs, .btn-group-xs > .btn": {
        padding: "1px 5px",
        fontSize: 12,
        lineHeight: "1.5",
        borderRadius: 3
      },
      ".btn-block": {
        display: "block",
        width: "100%"
      },
      ".btn-block + .btn-block": {
        marginTop: 5
      },
      'input[type="submit"].btn-block, input[type="reset"].btn-block, input[type="button"].btn-block': {
        width: "100%"
      },
      ".fade": {
        opacity: "0",
        W: "opacity .15s linear",
        O: "opacity .15s linear",
        transition: "opacity .15s linear"
      },
      ".fade.in": {
        opacity: "1"
      },
      ".collapse": {
        display: "none"
      },
      ".collapse.in": {
        display: "block"
      },
      "tr.collapse.in": {
        display: "table-row"
      },
      "tbody.collapse.in": {
        display: "table-row-group"
      },
      ".collapsing": {
        position: "relative",
        height: "0",
        overflow: "hidden",
        W: "height, visibility",
        O: "height, visibility",
        transitionTimingFunction: "ease",
        fallbacks: [
          {
            O: ".35s"
          },
          {
            W: ".35s"
          },
          {
            O: "ease"
          },
          {
            W: "ease"
          }
        ],
        transitionDuration: ".35s",
        transitionProperty: "height, visibility"
      },
      ".caret": {
        display: "inline-block",
        width: "0",
        height: "0",
        marginLeft: 2,
        verticalAlign: "middle",
        borderTop: "4px solid 9",
        fallbacks: [
          {
            borderTop: "4px dashed"
          }
        ],
        borderRight: "4px solid transparent",
        borderLeft: "4px solid transparent"
      },
      ".dropup, .dropdown": {
        position: "relative"
      },
      ".dropdown-toggle:focus": {
        outline: "0"
      },
      ".dropdown-menu": {
        position: "absolute",
        top: "100%",
        left: "0",
        zIndex: "1000",
        display: "none",
        float: "left",
        minWidth: 160,
        padding: "5px 0",
        margin: "2px 0 0",
        fontSize: 14,
        textAlign: "left",
        listStyle: "none",
        backgroundColor: "#fff",
        W: "0 6px 12px rgba(0, 0, 0, .175)",
        backgroundClip: "padding-box",
        border: "1px solid rgba(0, 0, 0, .15)",
        fallbacks: [
          {
            W: "padding-box"
          },
          {
            border: "1px solid #ccc"
          }
        ],
        borderRadius: 4,
        boxShadow: "0 6px 12px rgba(0, 0, 0, .175)"
      },
      ".dropdown-menu.pull-right": {
        right: "0",
        left: "auto"
      },
      ".dropdown-menu .divider": {
        height: 1,
        margin: "9px 0",
        overflow: "hidden",
        backgroundColor: "#e5e5e5"
      },
      ".dropdown-menu > li > a": {
        display: "block",
        padding: "3px 20px",
        clear: "both",
        fontWeight: "normal",
        lineHeight: "1.42857143",
        color: "#333",
        whiteSpace: "nowrap"
      },
      ".dropdown-menu > li > a:hover, .dropdown-menu > li > a:focus": {
        color: "#262626",
        textDecoration: "none",
        backgroundColor: "#f5f5f5"
      },
      ".dropdown-menu > .active > a, .dropdown-menu > .active > a:hover, .dropdown-menu > .active > a:focus": {
        color: "#fff",
        textDecoration: "none",
        backgroundColor: theme.colors.primary.main,
        outline: "0"
      },
      ".dropdown-menu > .disabled > a, .dropdown-menu > .disabled > a:hover, .dropdown-menu > .disabled > a:focus": {
        color: "#777"
      },
      ".dropdown-menu > .disabled > a:hover, .dropdown-menu > .disabled > a:focus": {
        textDecoration: "none",
        cursor: "not-allowed",
        backgroundColor: "transparent",
        backgroundImage: "none",
        filter: "progid:DXImageTransform.Microsoft.gradient(enabled = false)"
      },
      ".open > .dropdown-menu": {
        display: "block"
      },
      ".open > a": {
        outline: "0"
      },
      ".dropdown-menu-right": {
        right: "0",
        left: "auto"
      },
      ".dropdown-menu-left": {
        right: "auto",
        left: "0"
      },
      ".dropdown-header": {
        display: "block",
        padding: "3px 20px",
        fontSize: 12,
        lineHeight: "1.42857143",
        color: "#777",
        whiteSpace: "nowrap"
      },
      ".dropdown-backdrop": {
        position: "fixed",
        top: "0",
        right: "0",
        bottom: "0",
        left: "0",
        zIndex: "990"
      },
      ".pull-right > .dropdown-menu": {
        right: "0",
        left: "auto"
      },
      ".dropup .caret, .navbar-fixed-bottom .dropdown .caret": {
        content: '""',
        borderTop: "0",
        borderBottom: "4px solid 9",
        fallbacks: [
          {
            borderBottom: "4px dashed"
          }
        ]
      },
      ".dropup .dropdown-menu, .navbar-fixed-bottom .dropdown .dropdown-menu": {
        top: "auto",
        bottom: "100%",
        marginBottom: 2
      },
      ".btn-group, .btn-group-vertical": {
        position: "relative",
        display: "inline-block",
        verticalAlign: "middle"
      },
      ".btn-group > .btn, .btn-group-vertical > .btn": {
        position: "relative",
        float: "left"
      },
      ".btn-group > .btn:hover, .btn-group-vertical > .btn:hover, .btn-group > .btn:focus, .btn-group-vertical > .btn:focus, .btn-group > .btn:active, .btn-group-vertical > .btn:active, .btn-group > .btn.active, .btn-group-vertical > .btn.active": {
        zIndex: "2"
      },
      ".btn-group .btn + .btn, .btn-group .btn + .btn-group, .btn-group .btn-group + .btn, .btn-group .btn-group + .btn-group": {
        marginLeft: -1
      },
      ".btn-toolbar": {
        marginLeft: -5
      },
      ".btn-toolbar .btn, .btn-toolbar .btn-group, .btn-toolbar .input-group": {
        float: "left"
      },
      ".btn-toolbar > .btn, .btn-toolbar > .btn-group, .btn-toolbar > .input-group": {
        marginLeft: 5
      },
      ".btn-group > .btn:not(:first-child):not(:last-child):not(.dropdown-toggle)": {
        borderRadius: "0"
      },
      ".btn-group > .btn:first-child": {
        marginLeft: "0"
      },
      ".btn-group > .btn:first-child:not(:last-child):not(.dropdown-toggle)": {
        borderTopRightRadius: "0",
        borderBottomRightRadius: "0"
      },
      ".btn-group > .btn:last-child:not(:first-child), .btn-group > .dropdown-toggle:not(:first-child)": {
        borderTopLeftRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".btn-group > .btn-group": {
        float: "left"
      },
      ".btn-group > .btn-group:not(:first-child):not(:last-child) > .btn": {
        borderRadius: "0"
      },
      ".btn-group > .btn-group:first-child:not(:last-child) > .btn:last-child, .btn-group > .btn-group:first-child:not(:last-child) > .dropdown-toggle": {
        borderTopRightRadius: "0",
        borderBottomRightRadius: "0"
      },
      ".btn-group > .btn-group:last-child:not(:first-child) > .btn:first-child": {
        borderTopLeftRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".btn-group .dropdown-toggle:active, .btn-group.open .dropdown-toggle": {
        outline: "0"
      },
      ".btn-group > .btn + .dropdown-toggle": {
        paddingRight: 8,
        paddingLeft: 8
      },
      ".btn-group > .btn-lg + .dropdown-toggle": {
        paddingRight: 12,
        paddingLeft: 12
      },
      ".btn-group.open .dropdown-toggle": {
        W: "inset 0 3px 5px rgba(0, 0, 0, .125)",
        boxShadow: "inset 0 3px 5px rgba(0, 0, 0, .125)"
      },
      ".btn-group.open .dropdown-toggle.btn-link": {
        W: "none",
        boxShadow: "none"
      },
      ".btn .caret": {
        marginLeft: "0"
      },
      ".btn-lg .caret": {
        borderWidth: "5px 5px 0",
        borderBottomWidth: "0"
      },
      ".dropup .btn-lg .caret": {
        borderWidth: "0 5px 5px"
      },
      ".btn-group-vertical > .btn, .btn-group-vertical > .btn-group, .btn-group-vertical > .btn-group > .btn": {
        display: "block",
        float: "none",
        width: "100%",
        maxWidth: "100%"
      },
      ".btn-group-vertical > .btn-group > .btn": {
        float: "none"
      },
      ".btn-group-vertical > .btn + .btn, .btn-group-vertical > .btn + .btn-group, .btn-group-vertical > .btn-group + .btn, .btn-group-vertical > .btn-group + .btn-group": {
        marginTop: -1,
        marginLeft: "0"
      },
      ".btn-group-vertical > .btn:not(:first-child):not(:last-child)": {
        borderRadius: "0"
      },
      ".btn-group-vertical > .btn:first-child:not(:last-child)": {
        borderTopLeftRadius: 4,
        borderTopRightRadius: 4,
        borderBottomRightRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".btn-group-vertical > .btn:last-child:not(:first-child)": {
        borderTopLeftRadius: "0",
        borderTopRightRadius: "0",
        borderBottomRightRadius: 4,
        borderBottomLeftRadius: 4
      },
      ".btn-group-vertical > .btn-group:not(:first-child):not(:last-child) > .btn": {
        borderRadius: "0"
      },
      ".btn-group-vertical > .btn-group:first-child:not(:last-child) > .btn:last-child, .btn-group-vertical > .btn-group:first-child:not(:last-child) > .dropdown-toggle": {
        borderBottomRightRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".btn-group-vertical > .btn-group:last-child:not(:first-child) > .btn:first-child": {
        borderTopLeftRadius: "0",
        borderTopRightRadius: "0"
      },
      ".btn-group-justified": {
        display: "table",
        width: "100%",
        tableLayout: "fixed",
        borderCollapse: "separate"
      },
      ".btn-group-justified > .btn, .btn-group-justified > .btn-group": {
        display: "table-cell",
        float: "none",
        width: "1%"
      },
      ".btn-group-justified > .btn-group .btn": {
        width: "100%"
      },
      ".btn-group-justified > .btn-group .dropdown-menu": {
        left: "auto"
      },
      '[data-toggle="buttons"] > .btn input[type="radio"], [data-toggle="buttons"] > .btn-group > .btn input[type="radio"], [data-toggle="buttons"] > .btn input[type="checkbox"], [data-toggle="buttons"] > .btn-group > .btn input[type="checkbox"]': {
        position: "absolute",
        clip: "rect(0, 0, 0, 0)",
        pointerEvents: "none"
      },
      ".input-group": {
        position: "relative",
        display: "table",
        borderCollapse: "separate"
      },
      '.input-group[class*="col-"]': {
        float: "none",
        paddingRight: "0",
        paddingLeft: "0"
      },
      ".input-group .form-control": {
        position: "relative",
        zIndex: "2",
        float: "left",
        width: "100%",
        marginBottom: "0"
      },
      ".input-group .form-control:focus": {
        zIndex: "3"
      },
      ".input-group-lg > .form-control, .input-group-lg > .input-group-addon, .input-group-lg > .input-group-btn > .btn": {
        height: 46,
        padding: "10px 16px",
        fontSize: 18,
        lineHeight: "1.3333333",
        borderRadius: 6
      },
      "select.input-group-lg > .form-control, select.input-group-lg > .input-group-addon, select.input-group-lg > .input-group-btn > .btn": {
        height: 46,
        lineHeight: 46
      },
      "textarea.input-group-lg > .form-control, textarea.input-group-lg > .input-group-addon, textarea.input-group-lg > .input-group-btn > .btn, select[multiple].input-group-lg > .form-control, select[multiple].input-group-lg > .input-group-addon, select[multiple].input-group-lg > .input-group-btn > .btn": {
        height: "auto"
      },
      ".input-group-sm > .form-control, .input-group-sm > .input-group-addon, .input-group-sm > .input-group-btn > .btn": {
        height: 30,
        padding: "5px 10px",
        fontSize: 12,
        lineHeight: "1.5",
        borderRadius: 3
      },
      "select.input-group-sm > .form-control, select.input-group-sm > .input-group-addon, select.input-group-sm > .input-group-btn > .btn": {
        height: 30,
        lineHeight: 30
      },
      "textarea.input-group-sm > .form-control, textarea.input-group-sm > .input-group-addon, textarea.input-group-sm > .input-group-btn > .btn, select[multiple].input-group-sm > .form-control, select[multiple].input-group-sm > .input-group-addon, select[multiple].input-group-sm > .input-group-btn > .btn": {
        height: "auto"
      },
      ".input-group-addon, .input-group-btn, .input-group .form-control": {
        display: "table-cell"
      },
      ".input-group-addon:not(:first-child):not(:last-child), .input-group-btn:not(:first-child):not(:last-child), .input-group .form-control:not(:first-child):not(:last-child)": {
        borderRadius: "0"
      },
      ".input-group-addon, .input-group-btn": {
        width: "1%",
        whiteSpace: "nowrap",
        verticalAlign: "middle"
      },
      ".input-group-addon": {
        padding: "6px 12px",
        fontSize: 14,
        fontWeight: "normal",
        lineHeight: "1",
        color: "#555",
        textAlign: "center",
        backgroundColor: "#eee",
        border: "1px solid #ccc",
        borderRadius: 4
      },
      ".input-group-addon.input-sm": {
        padding: "5px 10px",
        fontSize: 12,
        borderRadius: 3
      },
      ".input-group-addon.input-lg": {
        padding: "10px 16px",
        fontSize: 18,
        borderRadius: 6
      },
      '.input-group-addon input[type="radio"], .input-group-addon input[type="checkbox"]': {
        marginTop: "0"
      },
      ".input-group .form-control:first-child, .input-group-addon:first-child, .input-group-btn:first-child > .btn, .input-group-btn:first-child > .btn-group > .btn, .input-group-btn:first-child > .dropdown-toggle, .input-group-btn:last-child > .btn:not(:last-child):not(.dropdown-toggle), .input-group-btn:last-child > .btn-group:not(:last-child) > .btn": {
        borderTopRightRadius: "0",
        borderBottomRightRadius: "0"
      },
      ".input-group-addon:first-child": {
        borderRight: "0"
      },
      ".input-group .form-control:last-child, .input-group-addon:last-child, .input-group-btn:last-child > .btn, .input-group-btn:last-child > .btn-group > .btn, .input-group-btn:last-child > .dropdown-toggle, .input-group-btn:first-child > .btn:not(:first-child), .input-group-btn:first-child > .btn-group:not(:first-child) > .btn": {
        borderTopLeftRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".input-group-addon:last-child": {
        borderLeft: "0"
      },
      ".input-group-btn": {
        position: "relative",
        fontSize: "0",
        whiteSpace: "nowrap"
      },
      ".input-group-btn > .btn": {
        position: "relative"
      },
      ".input-group-btn > .btn + .btn": {
        marginLeft: -1
      },
      ".input-group-btn > .btn:hover, .input-group-btn > .btn:focus, .input-group-btn > .btn:active": {
        zIndex: "2"
      },
      ".input-group-btn:first-child > .btn, .input-group-btn:first-child > .btn-group": {
        marginRight: -1
      },
      ".input-group-btn:last-child > .btn, .input-group-btn:last-child > .btn-group": {
        zIndex: "2",
        marginLeft: -1
      },
      ".nav": {
        paddingLeft: "0",
        marginBottom: "0",
        listStyle: "none"
      },
      ".nav > li": {
        position: "relative",
        display: "block"
      },
      ".nav > li > a": {
        position: "relative",
        display: "block",
        padding: "10px 15px"
      },
      ".nav > li > a:hover, .nav > li > a:focus": {
        textDecoration: "none",
        backgroundColor: "#eee"
      },
      ".nav > li.disabled > a": {
        color: "#777"
      },
      ".nav > li.disabled > a:hover, .nav > li.disabled > a:focus": {
        color: "#777",
        textDecoration: "none",
        cursor: "not-allowed",
        backgroundColor: "transparent"
      },
      ".nav .open > a, .nav .open > a:hover, .nav .open > a:focus": {
        backgroundColor: "#eee",
        borderColor: theme.colors.primary.main
      },
      ".nav .nav-divider": {
        height: 1,
        margin: "9px 0",
        overflow: "hidden",
        backgroundColor: "#e5e5e5"
      },
      ".nav > li > a > img": {
        maxWidth: "none"
      },
      ".nav-tabs": {
        borderBottom: "1px solid #ddd"
      },
      ".nav-tabs > li": {
        float: "left",
        marginBottom: -1
      },
      ".nav-tabs > li > a": {
        marginRight: 2,
        lineHeight: "1.42857143",
        border: "1px solid transparent",
        borderRadius: "4px 4px 0 0"
      },
      ".nav-tabs > li > a:hover": {
        borderColor: "#eee #eee #ddd"
      },
      ".nav-tabs > li.active > a, .nav-tabs > li.active > a:hover, .nav-tabs > li.active > a:focus": {
        color: "#555",
        cursor: "default",
        backgroundColor: "#fff",
        border: "1px solid #ddd",
        borderBottomColor: "transparent"
      },
      ".nav-tabs.nav-justified": {
        width: "100%",
        borderBottom: "0"
      },
      ".nav-tabs.nav-justified > li": {
        float: "none"
      },
      ".nav-tabs.nav-justified > li > a": {
        marginBottom: 5,
        textAlign: "center",
        marginRight: "0",
        borderRadius: 4
      },
      ".nav-tabs.nav-justified > .dropdown .dropdown-menu": {
        top: "auto",
        left: "auto"
      },
      ".nav-tabs.nav-justified > .active > a, .nav-tabs.nav-justified > .active > a:hover, .nav-tabs.nav-justified > .active > a:focus": {
        border: "1px solid #ddd"
      },
      ".nav-pills > li": {
        float: "left"
      },
      ".nav-pills > li > a": {
        borderRadius: 4
      },
      ".nav-pills > li + li": {
        marginLeft: 2
      },
      ".nav-pills > li.active > a, .nav-pills > li.active > a:hover, .nav-pills > li.active > a:focus": {
        color: "#fff",
        backgroundColor: theme.colors.primary.main
      },
      ".nav-stacked > li": {
        float: "none"
      },
      ".nav-stacked > li + li": {
        marginTop: 2,
        marginLeft: "0"
      },
      ".nav-justified": {
        width: "100%"
      },
      ".nav-justified > li": {
        float: "none"
      },
      ".nav-justified > li > a": {
        marginBottom: 5,
        textAlign: "center"
      },
      ".nav-justified > .dropdown .dropdown-menu": {
        top: "auto",
        left: "auto"
      },
      ".nav-tabs-justified": {
        borderBottom: "0"
      },
      ".nav-tabs-justified > li > a": {
        marginRight: "0",
        borderRadius: 4
      },
      ".nav-tabs-justified > .active > a, .nav-tabs-justified > .active > a:hover, .nav-tabs-justified > .active > a:focus": {
        border: "1px solid #ddd"
      },
      ".tab-content > .tab-pane": {
        display: "none"
      },
      ".tab-content > .active": {
        display: "block"
      },
      ".nav-tabs .dropdown-menu": {
        marginTop: -1,
        borderTopLeftRadius: "0",
        borderTopRightRadius: "0"
      },
      ".navbar": {
        position: "relative",
        minHeight: 50,
        marginBottom: 20,
        border: "1px solid transparent"
      },
      ".navbar-collapse": {
        paddingRight: 15,
        paddingLeft: 15,
        overflowX: "visible",
        W: "inset 0 1px 0 rgba(255, 255, 255, .1)",
        borderTop: "1px solid transparent",
        fallbacks: [
          {
            W: "touch"
          }
        ],
        boxShadow: "inset 0 1px 0 rgba(255, 255, 255, .1)"
      },
      ".navbar-collapse.in": {
        overflowY: "auto"
      },
      ".navbar-fixed-top .navbar-collapse, .navbar-fixed-bottom .navbar-collapse": {
        maxHeight: 340
      },
      "@media (max-device-width: 480px) and (orientation: landscape)": {
        ".navbar-fixed-top .navbar-collapse, .navbar-fixed-bottom .navbar-collapse": {
          maxHeight: 200
        }
      },
      ".container > .navbar-header, .container-fluid > .navbar-header, .container > .navbar-collapse, .container-fluid > .navbar-collapse": {
        marginRight: -15,
        marginLeft: -15
      },
      ".navbar-static-top": {
        zIndex: "1000",
        borderWidth: "0 0 1px"
      },
      ".navbar-fixed-top, .navbar-fixed-bottom": {
        position: "fixed",
        right: "0",
        left: "0",
        zIndex: "1030"
      },
      ".navbar-fixed-top": {
        top: "0",
        borderWidth: "0 0 1px"
      },
      ".navbar-fixed-bottom": {
        bottom: "0",
        marginBottom: "0",
        borderWidth: "1px 0 0"
      },
      ".navbar-brand": {
        float: "left",
        height: 50,
        padding: "15px 15px",
        fontSize: 18,
        lineHeight: 20
      },
      ".navbar-brand:hover, .navbar-brand:focus": {
        textDecoration: "none"
      },
      ".navbar-brand > img": {
        display: "block"
      },
      ".navbar-toggle": {
        position: "relative",
        float: "right",
        padding: "9px 10px",
        marginTop: 8,
        marginRight: 15,
        marginBottom: 8,
        backgroundColor: "transparent",
        backgroundImage: "none",
        border: "1px solid transparent",
        borderRadius: 4
      },
      ".navbar-toggle:focus": {
        outline: "0"
      },
      ".navbar-toggle .icon-bar": {
        display: "block",
        width: 22,
        height: 2,
        borderRadius: 1
      },
      ".navbar-toggle .icon-bar + .icon-bar": {
        marginTop: 4
      },
      ".navbar-nav": {
        margin: "7.5px -15px"
      },
      ".navbar-nav > li > a": {
        paddingTop: 10,
        paddingBottom: 10,
        lineHeight: 20
      },
      "@media (max-width: 767px)": {
        ".navbar-nav .open .dropdown-menu": {
          position: "static",
          float: "none",
          width: "auto",
          marginTop: "0",
          backgroundColor: "transparent",
          border: "0",
          W: "none",
          boxShadow: "none"
        },
        ".navbar-nav .open .dropdown-menu > li > a, .navbar-nav .open .dropdown-menu .dropdown-header": {
          padding: "5px 15px 5px 25px"
        },
        ".navbar-nav .open .dropdown-menu > li > a": {
          lineHeight: 20
        },
        ".navbar-nav .open .dropdown-menu > li > a:hover, .navbar-nav .open .dropdown-menu > li > a:focus": {
          backgroundImage: "none"
        },
        ".navbar-form .form-group": {
          marginBottom: 5
        },
        ".navbar-form .form-group:last-child": {
          marginBottom: "0"
        },
        ".navbar-default .navbar-nav .open .dropdown-menu > li > a": {
          color: "#777"
        },
        ".navbar-default .navbar-nav .open .dropdown-menu > li > a:hover, .navbar-default .navbar-nav .open .dropdown-menu > li > a:focus": {
          color: "#333",
          backgroundColor: "transparent"
        },
        ".navbar-default .navbar-nav .open .dropdown-menu > .active > a, .navbar-default .navbar-nav .open .dropdown-menu > .active > a:hover, .navbar-default .navbar-nav .open .dropdown-menu > .active > a:focus": {
          color: "#555",
          backgroundColor: "#e7e7e7"
        },
        ".navbar-default .navbar-nav .open .dropdown-menu > .disabled > a, .navbar-default .navbar-nav .open .dropdown-menu > .disabled > a:hover, .navbar-default .navbar-nav .open .dropdown-menu > .disabled > a:focus": {
          color: "#ccc",
          backgroundColor: "transparent"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu > .dropdown-header": {
          borderColor: "#080808"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu .divider": {
          backgroundColor: "#080808"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu > li > a": {
          color: "#9d9d9d"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu > li > a:hover, .navbar-inverse .navbar-nav .open .dropdown-menu > li > a:focus": {
          color: "#fff",
          backgroundColor: "transparent"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu > .active > a, .navbar-inverse .navbar-nav .open .dropdown-menu > .active > a:hover, .navbar-inverse .navbar-nav .open .dropdown-menu > .active > a:focus": {
          color: "#fff",
          backgroundColor: "#080808"
        },
        ".navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a, .navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a:hover, .navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a:focus": {
          color: "#444",
          backgroundColor: "transparent"
        },
        ".visible-xs": {
          display: "block !important"
        },
        "table.visible-xs": {
          display: "table !important"
        },
        "tr.visible-xs": {
          display: "table-row !important"
        },
        "th.visible-xs, td.visible-xs": {
          display: "table-cell !important"
        },
        ".visible-xs-block": {
          display: "block !important"
        },
        ".visible-xs-inline": {
          display: "inline !important"
        },
        ".visible-xs-inline-block": {
          display: "inline-block !important"
        },
        ".hidden-xs": {
          display: "none !important"
        }
      },
      ".navbar-form": {
        padding: "10px 15px",
        marginTop: 8,
        marginRight: -15,
        marginBottom: 8,
        marginLeft: -15,
        borderTop: "1px solid transparent",
        borderBottom: "1px solid transparent",
        W:
          "inset 0 1px 0 rgba(255, 255, 255, .1), 0 1px 0 rgba(255, 255, 255, .1)",
        boxShadow:
          "inset 0 1px 0 rgba(255, 255, 255, .1), 0 1px 0 rgba(255, 255, 255, .1)"
      },
      ".navbar-nav > li > .dropdown-menu": {
        marginTop: "0",
        borderTopLeftRadius: "0",
        borderTopRightRadius: "0"
      },
      ".navbar-fixed-bottom .navbar-nav > li > .dropdown-menu": {
        marginBottom: "0",
        borderTopLeftRadius: 4,
        borderTopRightRadius: 4,
        borderBottomRightRadius: "0",
        borderBottomLeftRadius: "0"
      },
      ".navbar-btn": {
        marginTop: 8,
        marginBottom: 8
      },
      ".navbar-btn.btn-sm": {
        marginTop: 10,
        marginBottom: 10
      },
      ".navbar-btn.btn-xs": {
        marginTop: 14,
        marginBottom: 14
      },
      ".navbar-text": {
        marginTop: 15,
        marginBottom: 15
      },
      ".navbar-default": {
        backgroundColor: "#f8f8f8",
        borderColor: "#e7e7e7"
      },
      ".navbar-default .navbar-brand": {
        color: "#777"
      },
      ".navbar-default .navbar-brand:hover, .navbar-default .navbar-brand:focus": {
        color: "#5e5e5e",
        backgroundColor: "transparent"
      },
      ".navbar-default .navbar-text": {
        color: "#777"
      },
      ".navbar-default .navbar-nav > li > a": {
        color: "#777"
      },
      ".navbar-default .navbar-nav > li > a:hover, .navbar-default .navbar-nav > li > a:focus": {
        color: "#333",
        backgroundColor: "transparent"
      },
      ".navbar-default .navbar-nav > .active > a, .navbar-default .navbar-nav > .active > a:hover, .navbar-default .navbar-nav > .active > a:focus": {
        color: "#555",
        backgroundColor: "#e7e7e7"
      },
      ".navbar-default .navbar-nav > .disabled > a, .navbar-default .navbar-nav > .disabled > a:hover, .navbar-default .navbar-nav > .disabled > a:focus": {
        color: "#ccc",
        backgroundColor: "transparent"
      },
      ".navbar-default .navbar-toggle": {
        borderColor: "#ddd"
      },
      ".navbar-default .navbar-toggle:hover, .navbar-default .navbar-toggle:focus": {
        backgroundColor: "#ddd"
      },
      ".navbar-default .navbar-toggle .icon-bar": {
        backgroundColor: "#888"
      },
      ".navbar-default .navbar-collapse, .navbar-default .navbar-form": {
        borderColor: "#e7e7e7"
      },
      ".navbar-default .navbar-nav > .open > a, .navbar-default .navbar-nav > .open > a:hover, .navbar-default .navbar-nav > .open > a:focus": {
        color: "#555",
        backgroundColor: "#e7e7e7"
      },
      ".navbar-default .navbar-link": {
        color: "#777"
      },
      ".navbar-default .navbar-link:hover": {
        color: "#333"
      },
      ".navbar-default .btn-link": {
        color: "#777"
      },
      ".navbar-default .btn-link:hover, .navbar-default .btn-link:focus": {
        color: "#333"
      },
      ".navbar-default .btn-link[disabled]:hover, fieldset[disabled] .navbar-default .btn-link:hover, .navbar-default .btn-link[disabled]:focus, fieldset[disabled] .navbar-default .btn-link:focus": {
        color: "#ccc"
      },
      ".navbar-inverse": {
        backgroundColor: "#222",
        borderColor: "#080808"
      },
      ".navbar-inverse .navbar-brand": {
        color: "#9d9d9d"
      },
      ".navbar-inverse .navbar-brand:hover, .navbar-inverse .navbar-brand:focus": {
        color: "#fff",
        backgroundColor: "transparent"
      },
      ".navbar-inverse .navbar-text": {
        color: "#9d9d9d"
      },
      ".navbar-inverse .navbar-nav > li > a": {
        color: "#9d9d9d"
      },
      ".navbar-inverse .navbar-nav > li > a:hover, .navbar-inverse .navbar-nav > li > a:focus": {
        color: "#fff",
        backgroundColor: "transparent"
      },
      ".navbar-inverse .navbar-nav > .active > a, .navbar-inverse .navbar-nav > .active > a:hover, .navbar-inverse .navbar-nav > .active > a:focus": {
        color: "#fff",
        backgroundColor: "#080808"
      },
      ".navbar-inverse .navbar-nav > .disabled > a, .navbar-inverse .navbar-nav > .disabled > a:hover, .navbar-inverse .navbar-nav > .disabled > a:focus": {
        color: "#444",
        backgroundColor: "transparent"
      },
      ".navbar-inverse .navbar-toggle": {
        borderColor: "#333"
      },
      ".navbar-inverse .navbar-toggle:hover, .navbar-inverse .navbar-toggle:focus": {
        backgroundColor: "#333"
      },
      ".navbar-inverse .navbar-toggle .icon-bar": {
        backgroundColor: "#fff"
      },
      ".navbar-inverse .navbar-collapse, .navbar-inverse .navbar-form": {
        borderColor: "#101010"
      },
      ".navbar-inverse .navbar-nav > .open > a, .navbar-inverse .navbar-nav > .open > a:hover, .navbar-inverse .navbar-nav > .open > a:focus": {
        color: "#fff",
        backgroundColor: "#080808"
      },
      ".navbar-inverse .navbar-link": {
        color: "#9d9d9d"
      },
      ".navbar-inverse .navbar-link:hover": {
        color: "#fff"
      },
      ".navbar-inverse .btn-link": {
        color: "#9d9d9d"
      },
      ".navbar-inverse .btn-link:hover, .navbar-inverse .btn-link:focus": {
        color: "#fff"
      },
      ".navbar-inverse .btn-link[disabled]:hover, fieldset[disabled] .navbar-inverse .btn-link:hover, .navbar-inverse .btn-link[disabled]:focus, fieldset[disabled] .navbar-inverse .btn-link:focus": {
        color: "#444"
      },
      ".breadcrumb": {
        padding: "8px 15px",
        marginBottom: 20,
        listStyle: "none",
        backgroundColor: "#f5f5f5",
        borderRadius: 4
      },
      ".breadcrumb > li": {
        display: "inline-block"
      },
      ".breadcrumb > li + li:before": {
        padding: "0 5px",
        color: "#ccc",
        content: '"/\\00a0"'
      },
      ".breadcrumb > .active": {
        color: "#777"
      },
      ".pagination": {
        display: "inline-block",
        paddingLeft: "0",
        margin: "20px 0",
        borderRadius: 4
      },
      ".pagination > li": {
        display: "inline"
      },
      ".pagination > li > a, .pagination > li > span": {
        position: "relative",
        float: "left",
        padding: "6px 12px",
        marginLeft: -1,
        lineHeight: "1.42857143",
        color: theme.colors.primary.main,
        textDecoration: "none",
        backgroundColor: "#fff",
        border: "1px solid #ddd"
      },
      ".pagination > li:first-child > a, .pagination > li:first-child > span": {
        marginLeft: "0",
        borderTopLeftRadius: 4,
        borderBottomLeftRadius: 4
      },
      ".pagination > li:last-child > a, .pagination > li:last-child > span": {
        borderTopRightRadius: 4,
        borderBottomRightRadius: 4
      },
      ".pagination > li > a:hover, .pagination > li > span:hover, .pagination > li > a:focus, .pagination > li > span:focus": {
        zIndex: "2",
        color: "#23527c",
        backgroundColor: "#eee",
        borderColor: "#ddd"
      },
      ".pagination > .active > a, .pagination > .active > span, .pagination > .active > a:hover, .pagination > .active > span:hover, .pagination > .active > a:focus, .pagination > .active > span:focus": {
        zIndex: "3",
        color: "#fff",
        cursor: "default",
        backgroundColor: theme.colors.primary.main,
        borderColor: theme.colors.primary.main
      },
      ".pagination > .disabled > span, .pagination > .disabled > span:hover, .pagination > .disabled > span:focus, .pagination > .disabled > a, .pagination > .disabled > a:hover, .pagination > .disabled > a:focus": {
        color: "#777",
        cursor: "not-allowed",
        backgroundColor: "#fff",
        borderColor: "#ddd"
      },
      ".pagination-lg > li > a, .pagination-lg > li > span": {
        padding: "10px 16px",
        fontSize: 18,
        lineHeight: "1.3333333"
      },
      ".pagination-lg > li:first-child > a, .pagination-lg > li:first-child > span": {
        borderTopLeftRadius: 6,
        borderBottomLeftRadius: 6
      },
      ".pagination-lg > li:last-child > a, .pagination-lg > li:last-child > span": {
        borderTopRightRadius: 6,
        borderBottomRightRadius: 6
      },
      ".pagination-sm > li > a, .pagination-sm > li > span": {
        padding: "5px 10px",
        fontSize: 12,
        lineHeight: "1.5"
      },
      ".pagination-sm > li:first-child > a, .pagination-sm > li:first-child > span": {
        borderTopLeftRadius: 3,
        borderBottomLeftRadius: 3
      },
      ".pagination-sm > li:last-child > a, .pagination-sm > li:last-child > span": {
        borderTopRightRadius: 3,
        borderBottomRightRadius: 3
      },
      ".pager": {
        paddingLeft: "0",
        margin: "20px 0",
        textAlign: "center",
        listStyle: "none"
      },
      ".pager li": {
        display: "inline"
      },
      ".pager li > a, .pager li > span": {
        display: "inline-block",
        padding: "5px 14px",
        backgroundColor: "#fff",
        border: "1px solid #ddd",
        borderRadius: 15
      },
      ".pager li > a:hover, .pager li > a:focus": {
        textDecoration: "none",
        backgroundColor: "#eee"
      },
      ".pager .next > a, .pager .next > span": {
        float: "right"
      },
      ".pager .previous > a, .pager .previous > span": {
        float: "left"
      },
      ".pager .disabled > a, .pager .disabled > a:hover, .pager .disabled > a:focus, .pager .disabled > span": {
        color: "#777",
        cursor: "not-allowed",
        backgroundColor: "#fff"
      },
      ".label": {
        ...theme.typography.caption,
        display: "inline",
        padding: ".2em .6em .3em",
        lineHeight: "1",
        color: "#fff",
        textAlign: "center",
        whiteSpace: "nowrap",
        verticalAlign: "baseline",
        borderRadius: ".25em"
      },
      "a.label:hover, a.label:focus": {
        color: "#fff",
        textDecoration: "none",
        cursor: "pointer"
      },
      ".label:empty": {
        display: "none"
      },
      ".btn .label": {
        position: "relative",
        top: -1
      },
      ".label-default": {
        backgroundColor: "#777"
      },
      ".label-default[href]:hover, .label-default[href]:focus": {
        backgroundColor: "#5e5e5e"
      },
      ".label-primary": {
        backgroundColor: theme.colors.primary.main
      },
      ".label-primary[href]:hover, .label-primary[href]:focus": {
        backgroundColor: theme.colors.primary.dark
      },
      ".label-success": {
        backgroundColor: theme.colors.success.main
      },
      ".label-success[href]:hover, .label-success[href]:focus": {
        backgroundColor: theme.colors.success.dark
      },
      ".label-info": {
        backgroundColor: "#5bc0de"
      },
      ".label-info[href]:hover, .label-info[href]:focus": {
        backgroundColor: "#31b0d5"
      },
      ".label-warning": {
        backgroundColor: "#f0ad4e"
      },
      ".label-warning[href]:hover, .label-warning[href]:focus": {
        backgroundColor: "#ec971f"
      },
      ".label-danger": {
        backgroundColor: theme.colors.error.main
      },
      ".label-danger[href]:hover, .label-danger[href]:focus": {
        backgroundColor: theme.colors.error.dark
      },
      ".badge": {
        display: "inline-block",
        minWidth: 10,
        padding: "3px 7px",
        fontSize: 12,
        fontWeight: "bold",
        lineHeight: "1",
        color: "#fff",
        textAlign: "center",
        whiteSpace: "nowrap",
        verticalAlign: "middle",
        backgroundColor: "#777",
        borderRadius: 10
      },
      ".badge:empty": {
        display: "none"
      },
      ".btn .badge": {
        position: "relative",
        top: -1
      },
      ".btn-xs .badge, .btn-group-xs > .btn .badge": {
        top: "0",
        padding: "1px 5px"
      },
      "a.badge:hover, a.badge:focus": {
        color: "#fff",
        textDecoration: "none",
        cursor: "pointer"
      },
      ".list-group-item.active > .badge, .nav-pills > .active > a > .badge": {
        color: theme.colors.primary.main,
        backgroundColor: "#fff"
      },
      ".list-group-item > .badge": {
        float: "right"
      },
      ".list-group-item > .badge + .badge": {
        marginRight: 5
      },
      ".nav-pills > li > a > .badge": {
        marginLeft: 3
      },
      ".jumbotron": {
        paddingTop: 30,
        paddingBottom: 30,
        marginBottom: 30,
        color: "inherit",
        backgroundColor: "#eee"
      },
      ".jumbotron h1, .jumbotron .h1": {
        color: "inherit"
      },
      ".jumbotron p": {
        marginBottom: 15,
        fontSize: 21,
        fontWeight: "200"
      },
      ".jumbotron > hr": {
        borderTopColor: "#d5d5d5"
      },
      ".container .jumbotron, .container-fluid .jumbotron": {
        paddingRight: 15,
        paddingLeft: 15,
        borderRadius: 6
      },
      ".jumbotron .container": {
        maxWidth: "100%"
      },
      "@media screen and (min-width: 768px)": {
        ".jumbotron": {
          paddingTop: 48,
          paddingBottom: 48
        },
        ".container .jumbotron, .container-fluid .jumbotron": {
          paddingRight: 60,
          paddingLeft: 60
        },
        ".jumbotron h1, .jumbotron .h1": {
          fontSize: 63
        },
        ".carousel-control .glyphicon-chevron-left, .carousel-control .glyphicon-chevron-right, .carousel-control .icon-prev, .carousel-control .icon-next": {
          width: 30,
          height: 30,
          marginTop: -10,
          fontSize: 30
        },
        ".carousel-control .glyphicon-chevron-left, .carousel-control .icon-prev": {
          marginLeft: -10
        },
        ".carousel-control .glyphicon-chevron-right, .carousel-control .icon-next": {
          marginRight: -10
        },
        ".carousel-caption": {
          right: "20%",
          left: "20%",
          paddingBottom: 30
        },
        ".carousel-indicators": {
          bottom: 20
        }
      },
      ".thumbnail": {
        display: "block",
        padding: 4,
        marginBottom: 20,
        lineHeight: "1.42857143",
        backgroundColor: "#fff",
        border: "1px solid #ddd",
        borderRadius: 4,
        W: "border .2s ease-in-out",
        O: "border .2s ease-in-out",
        transition: "border .2s ease-in-out"
      },
      ".thumbnail > img, .thumbnail a > img": {
        marginRight: "auto",
        marginLeft: "auto"
      },
      "a.thumbnail:hover, a.thumbnail:focus, a.thumbnail.active": {
        borderColor: theme.colors.primary.main
      },
      ".thumbnail .caption": {
        padding: 9,
        color: "#333"
      },
      ".alert": {
        padding: 15,
        marginBottom: 20,
        border: "1px solid transparent",
        borderRadius: 4
      },
      ".alert h4": {
        marginTop: "0",
        color: "inherit"
      },
      ".alert .alert-link": {
        fontWeight: "bold"
      },
      ".alert > p, .alert > ul": {
        marginBottom: "0"
      },
      ".alert > p + p": {
        marginTop: 5
      },
      ".alert-dismissable, .alert-dismissible": {
        paddingRight: 35
      },
      ".alert-dismissable .close, .alert-dismissible .close": {
        position: "relative",
        top: -2,
        right: -21,
        color: "inherit"
      },
      ".alert-success": {
        color: "#3c763d",
        backgroundColor: "#dff0d8",
        borderColor: "#d6e9c6"
      },
      ".alert-success hr": {
        borderTopColor: "#c9e2b3"
      },
      ".alert-success .alert-link": {
        color: "#2b542c"
      },
      ".alert-info": {
        color: "#31708f",
        backgroundColor: "#d9edf7",
        borderColor: "#bce8f1"
      },
      ".alert-info hr": {
        borderTopColor: "#a6e1ec"
      },
      ".alert-info .alert-link": {
        color: "#245269"
      },
      ".alert-warning": {
        color: "#8a6d3b",
        backgroundColor: "#fcf8e3",
        borderColor: "#faebcc"
      },
      ".alert-warning hr": {
        borderTopColor: "#f7e1b5"
      },
      ".alert-warning .alert-link": {
        color: "#66512c"
      },
      ".alert-danger": {
        color: "#a94442",
        backgroundColor: "#f2dede",
        borderColor: "#ebccd1"
      },
      ".alert-danger hr": {
        borderTopColor: "#e4b9c0"
      },
      ".alert-danger .alert-link": {
        color: "#843534"
      },
      "@keyframes progress-bar-stripes": {
        from: {
          backgroundPosition: "40px 0"
        },
        to: {
          backgroundPosition: "0 0"
        }
      },
      ".progress": {
        height: 20,
        marginBottom: 20,
        overflow: "hidden",
        backgroundColor: "#f5f5f5",
        borderRadius: 4,
        W: "inset 0 1px 2px rgba(0, 0, 0, .1)",
        boxShadow: "inset 0 1px 2px rgba(0, 0, 0, .1)"
      },
      ".progress-bar": {
        float: "left",
        width: "0",
        height: "100%",
        fontSize: 12,
        lineHeight: 20,
        color: "#fff",
        textAlign: "center",
        backgroundColor: theme.colors.primary.main,
        W: "width .6s ease",
        boxShadow: "inset 0 -1px 0 rgba(0, 0, 0, .15)",
        fallbacks: [
          {
            W: "inset 0 -1px 0 rgba(0, 0, 0, .15)"
          }
        ],
        O: "width .6s ease",
        transition: "width .6s ease"
      },
      ".progress-striped .progress-bar, .progress-bar-striped": {
        backgroundImage:
          "linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)",
        fallbacks: [
          {
            backgroundImage:
              "-o-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          }
        ],
        W: "40px 40px",
        backgroundSize: "40px 40px"
      },
      ".progress.active .progress-bar, .progress-bar.active": {
        W: "progress-bar-stripes 2s linear infinite",
        O: "progress-bar-stripes 2s linear infinite",
        animation: "progress-bar-stripes 2s linear infinite"
      },
      ".progress-bar-success": {
        backgroundColor: theme.colors.success.main
      },
      ".progress-striped .progress-bar-success": {
        backgroundImage:
          "linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)",
        fallbacks: [
          {
            backgroundImage:
              "-o-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          }
        ]
      },
      ".progress-bar-info": {
        backgroundColor: "#5bc0de"
      },
      ".progress-striped .progress-bar-info": {
        backgroundImage:
          "linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)",
        fallbacks: [
          {
            backgroundImage:
              "-o-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          }
        ]
      },
      ".progress-bar-warning": {
        backgroundColor: "#f0ad4e"
      },
      ".progress-striped .progress-bar-warning": {
        backgroundImage:
          "linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)",
        fallbacks: [
          {
            backgroundImage:
              "-o-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          }
        ]
      },
      ".progress-bar-danger": {
        backgroundColor: theme.colors.error.main
      },
      ".progress-striped .progress-bar-danger": {
        backgroundImage:
          "linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)",
        fallbacks: [
          {
            backgroundImage:
              "-o-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(45deg, rgba(255, 255, 255, .15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, .15) 50%, rgba(255, 255, 255, .15) 75%, transparent 75%, transparent)"
          }
        ]
      },
      ".media": {
        marginTop: 15
      },
      ".media:first-child": {
        marginTop: "0"
      },
      ".media, .media-body": {
        overflow: "hidden",
        zoom: "1"
      },
      ".media-body": {
        width: 10000
      },
      ".media-object": {
        display: "block"
      },
      ".media-object.img-thumbnail": {
        maxWidth: "none"
      },
      ".media-right, .media > .pull-right": {
        paddingLeft: 10
      },
      ".media-left, .media > .pull-left": {
        paddingRight: 10
      },
      ".media-left, .media-right, .media-body": {
        display: "table-cell",
        verticalAlign: "top"
      },
      ".media-middle": {
        verticalAlign: "middle"
      },
      ".media-bottom": {
        verticalAlign: "bottom"
      },
      ".media-heading": {
        marginTop: "0",
        marginBottom: 5
      },
      ".media-list": {
        paddingLeft: "0",
        listStyle: "none"
      },
      ".list-group": {
        paddingLeft: "0",
        marginBottom: 20
      },
      ".list-group-item": {
        position: "relative",
        display: "block",
        padding: "10px 15px",
        marginBottom: -1,
        backgroundColor: "#fff",
        border: "1px solid #ddd"
      },
      ".list-group-item:first-child": {
        borderTopLeftRadius: 4,
        borderTopRightRadius: 4
      },
      ".list-group-item:last-child": {
        marginBottom: "0",
        borderBottomRightRadius: 4,
        borderBottomLeftRadius: 4
      },
      "a.list-group-item, button.list-group-item": {
        color: "#555"
      },
      "a.list-group-item .list-group-item-heading, button.list-group-item .list-group-item-heading": {
        color: "#333"
      },
      "a.list-group-item:hover, button.list-group-item:hover, a.list-group-item:focus, button.list-group-item:focus": {
        color: "#555",
        textDecoration: "none",
        backgroundColor: "#f5f5f5"
      },
      "button.list-group-item": {
        width: "100%",
        textAlign: "left"
      },
      ".list-group-item.disabled, .list-group-item.disabled:hover, .list-group-item.disabled:focus": {
        color: "#777",
        cursor: "not-allowed",
        backgroundColor: "#eee"
      },
      ".list-group-item.disabled .list-group-item-heading, .list-group-item.disabled:hover .list-group-item-heading, .list-group-item.disabled:focus .list-group-item-heading": {
        color: "inherit"
      },
      ".list-group-item.disabled .list-group-item-text, .list-group-item.disabled:hover .list-group-item-text, .list-group-item.disabled:focus .list-group-item-text": {
        color: "#777"
      },
      ".list-group-item.active, .list-group-item.active:hover, .list-group-item.active:focus": {
        zIndex: "2",
        color: "#fff",
        backgroundColor: theme.colors.primary.main,
        borderColor: theme.colors.primary.main
      },
      ".list-group-item.active .list-group-item-heading, .list-group-item.active:hover .list-group-item-heading, .list-group-item.active:focus .list-group-item-heading, .list-group-item.active .list-group-item-heading > small, .list-group-item.active:hover .list-group-item-heading > small, .list-group-item.active:focus .list-group-item-heading > small, .list-group-item.active .list-group-item-heading > .small, .list-group-item.active:hover .list-group-item-heading > .small, .list-group-item.active:focus .list-group-item-heading > .small": {
        color: "inherit"
      },
      ".list-group-item.active .list-group-item-text, .list-group-item.active:hover .list-group-item-text, .list-group-item.active:focus .list-group-item-text": {
        color: "#c7ddef"
      },
      ".list-group-item-success": {
        color: "#3c763d",
        backgroundColor: "#dff0d8"
      },
      "a.list-group-item-success, button.list-group-item-success": {
        color: "#3c763d"
      },
      "a.list-group-item-success .list-group-item-heading, button.list-group-item-success .list-group-item-heading": {
        color: "inherit"
      },
      "a.list-group-item-success:hover, button.list-group-item-success:hover, a.list-group-item-success:focus, button.list-group-item-success:focus": {
        color: "#3c763d",
        backgroundColor: "#d0e9c6"
      },
      "a.list-group-item-success.active, button.list-group-item-success.active, a.list-group-item-success.active:hover, button.list-group-item-success.active:hover, a.list-group-item-success.active:focus, button.list-group-item-success.active:focus": {
        color: "#fff",
        backgroundColor: "#3c763d",
        borderColor: "#3c763d"
      },
      ".list-group-item-info": {
        color: "#31708f",
        backgroundColor: "#d9edf7"
      },
      "a.list-group-item-info, button.list-group-item-info": {
        color: "#31708f"
      },
      "a.list-group-item-info .list-group-item-heading, button.list-group-item-info .list-group-item-heading": {
        color: "inherit"
      },
      "a.list-group-item-info:hover, button.list-group-item-info:hover, a.list-group-item-info:focus, button.list-group-item-info:focus": {
        color: "#31708f",
        backgroundColor: "#c4e3f3"
      },
      "a.list-group-item-info.active, button.list-group-item-info.active, a.list-group-item-info.active:hover, button.list-group-item-info.active:hover, a.list-group-item-info.active:focus, button.list-group-item-info.active:focus": {
        color: "#fff",
        backgroundColor: "#31708f",
        borderColor: "#31708f"
      },
      ".list-group-item-warning": {
        color: "#8a6d3b",
        backgroundColor: "#fcf8e3"
      },
      "a.list-group-item-warning, button.list-group-item-warning": {
        color: "#8a6d3b"
      },
      "a.list-group-item-warning .list-group-item-heading, button.list-group-item-warning .list-group-item-heading": {
        color: "inherit"
      },
      "a.list-group-item-warning:hover, button.list-group-item-warning:hover, a.list-group-item-warning:focus, button.list-group-item-warning:focus": {
        color: "#8a6d3b",
        backgroundColor: "#faf2cc"
      },
      "a.list-group-item-warning.active, button.list-group-item-warning.active, a.list-group-item-warning.active:hover, button.list-group-item-warning.active:hover, a.list-group-item-warning.active:focus, button.list-group-item-warning.active:focus": {
        color: "#fff",
        backgroundColor: "#8a6d3b",
        borderColor: "#8a6d3b"
      },
      ".list-group-item-danger": {
        color: "#a94442",
        backgroundColor: "#f2dede"
      },
      "a.list-group-item-danger, button.list-group-item-danger": {
        color: "#a94442"
      },
      "a.list-group-item-danger .list-group-item-heading, button.list-group-item-danger .list-group-item-heading": {
        color: "inherit"
      },
      "a.list-group-item-danger:hover, button.list-group-item-danger:hover, a.list-group-item-danger:focus, button.list-group-item-danger:focus": {
        color: "#a94442",
        backgroundColor: "#ebcccc"
      },
      "a.list-group-item-danger.active, button.list-group-item-danger.active, a.list-group-item-danger.active:hover, button.list-group-item-danger.active:hover, a.list-group-item-danger.active:focus, button.list-group-item-danger.active:focus": {
        color: "#fff",
        backgroundColor: "#a94442",
        borderColor: "#a94442"
      },
      ".list-group-item-heading": {
        marginTop: "0",
        marginBottom: 5
      },
      ".list-group-item-text": {
        marginBottom: "0",
        lineHeight: "1.3"
      },
      ".panel": {
        marginBottom: 20,
        backgroundColor: "#fff",
        border: "1px solid #eee",
        borderRadius: theme.spacing / 2,
        boxShadow: "5px 5px 10px #0000000a"
      },
      ".panel-body": {
        padding: 15
      },
      ".panel-heading": {
        padding: "10px 15px",
        borderBottom: "1px solid transparent",
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3,
        display: "flex" as "flex",
        alignItems: "center" as "center"
      },
      ".panel-heading > .dropdown .dropdown-toggle": {
        color: "inherit"
      },
      ".panel-title": {
        "&:before": {
          content: "''",
          position: "absolute" as "absolute",
          width: 4 * theme.spacing,
          height: 2,
          background: theme.colors.primary.dark,
          bottom: -theme.spacing,
          left: 0
        },
        position: "relative" as "relative",
        flex: 1,
        marginTop: "0",
        marginBottom: "0",
        fontSize: 16,
        fontWeight: 600,
        color: "inherit"
      },
      ".panel-title > a, .panel-title > small, .panel-title > .small, .panel-title > small > a, .panel-title > .small > a": {
        color: "inherit"
      },
      ".panel-footer": {
        padding: "10px 15px",
        borderBottomRightRadius: 3,
        borderBottomLeftRadius: 3
      },
      ".panel > .list-group, .panel > .panel-collapse > .list-group": {
        marginBottom: "0"
      },
      ".panel > .list-group .list-group-item, .panel > .panel-collapse > .list-group .list-group-item": {
        borderWidth: "1px 0",
        borderRadius: "0"
      },
      ".panel > .list-group:first-child .list-group-item:first-child, .panel > .panel-collapse > .list-group:first-child .list-group-item:first-child": {
        borderTop: "0",
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3
      },
      ".panel > .list-group:last-child .list-group-item:last-child, .panel > .panel-collapse > .list-group:last-child .list-group-item:last-child": {
        borderBottom: "0",
        borderBottomRightRadius: 3,
        borderBottomLeftRadius: 3
      },
      ".panel > .panel-heading + .panel-collapse > .list-group .list-group-item:first-child": {
        borderTopLeftRadius: "0",
        borderTopRightRadius: "0"
      },
      ".panel-heading + .list-group .list-group-item:first-child": {
        borderTopWidth: "0"
      },
      ".list-group + .panel-footer": {
        borderTopWidth: "0"
      },
      ".panel > .table, .panel > .table-responsive > .table, .panel > .panel-collapse > .table": {
        marginBottom: "0"
      },
      ".panel > .table caption, .panel > .table-responsive > .table caption, .panel > .panel-collapse > .table caption": {
        paddingRight: 15,
        paddingLeft: 15
      },
      ".panel > .table:first-child, .panel > .table-responsive:first-child > .table:first-child": {
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3
      },
      ".panel > .table:first-child > thead:first-child > tr:first-child, .panel > .table-responsive:first-child > .table:first-child > thead:first-child > tr:first-child, .panel > .table:first-child > tbody:first-child > tr:first-child, .panel > .table-responsive:first-child > .table:first-child > tbody:first-child > tr:first-child": {
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3
      },
      ".panel > .table:first-child > thead:first-child > tr:first-child td:first-child, .panel > .table-responsive:first-child > .table:first-child > thead:first-child > tr:first-child td:first-child, .panel > .table:first-child > tbody:first-child > tr:first-child td:first-child, .panel > .table-responsive:first-child > .table:first-child > tbody:first-child > tr:first-child td:first-child, .panel > .table:first-child > thead:first-child > tr:first-child th:first-child, .panel > .table-responsive:first-child > .table:first-child > thead:first-child > tr:first-child th:first-child, .panel > .table:first-child > tbody:first-child > tr:first-child th:first-child, .panel > .table-responsive:first-child > .table:first-child > tbody:first-child > tr:first-child th:first-child": {
        borderTopLeftRadius: 3
      },
      ".panel > .table:first-child > thead:first-child > tr:first-child td:last-child, .panel > .table-responsive:first-child > .table:first-child > thead:first-child > tr:first-child td:last-child, .panel > .table:first-child > tbody:first-child > tr:first-child td:last-child, .panel > .table-responsive:first-child > .table:first-child > tbody:first-child > tr:first-child td:last-child, .panel > .table:first-child > thead:first-child > tr:first-child th:last-child, .panel > .table-responsive:first-child > .table:first-child > thead:first-child > tr:first-child th:last-child, .panel > .table:first-child > tbody:first-child > tr:first-child th:last-child, .panel > .table-responsive:first-child > .table:first-child > tbody:first-child > tr:first-child th:last-child": {
        borderTopRightRadius: 3
      },
      ".panel > .table:last-child, .panel > .table-responsive:last-child > .table:last-child": {
        borderBottomRightRadius: 3,
        borderBottomLeftRadius: 3
      },
      ".panel > .table:last-child > tbody:last-child > tr:last-child, .panel > .table-responsive:last-child > .table:last-child > tbody:last-child > tr:last-child, .panel > .table:last-child > tfoot:last-child > tr:last-child, .panel > .table-responsive:last-child > .table:last-child > tfoot:last-child > tr:last-child": {
        borderBottomRightRadius: 3,
        borderBottomLeftRadius: 3
      },
      ".panel > .table:last-child > tbody:last-child > tr:last-child td:first-child, .panel > .table-responsive:last-child > .table:last-child > tbody:last-child > tr:last-child td:first-child, .panel > .table:last-child > tfoot:last-child > tr:last-child td:first-child, .panel > .table-responsive:last-child > .table:last-child > tfoot:last-child > tr:last-child td:first-child, .panel > .table:last-child > tbody:last-child > tr:last-child th:first-child, .panel > .table-responsive:last-child > .table:last-child > tbody:last-child > tr:last-child th:first-child, .panel > .table:last-child > tfoot:last-child > tr:last-child th:first-child, .panel > .table-responsive:last-child > .table:last-child > tfoot:last-child > tr:last-child th:first-child": {
        borderBottomLeftRadius: 3
      },
      ".panel > .table:last-child > tbody:last-child > tr:last-child td:last-child, .panel > .table-responsive:last-child > .table:last-child > tbody:last-child > tr:last-child td:last-child, .panel > .table:last-child > tfoot:last-child > tr:last-child td:last-child, .panel > .table-responsive:last-child > .table:last-child > tfoot:last-child > tr:last-child td:last-child, .panel > .table:last-child > tbody:last-child > tr:last-child th:last-child, .panel > .table-responsive:last-child > .table:last-child > tbody:last-child > tr:last-child th:last-child, .panel > .table:last-child > tfoot:last-child > tr:last-child th:last-child, .panel > .table-responsive:last-child > .table:last-child > tfoot:last-child > tr:last-child th:last-child": {
        borderBottomRightRadius: 3
      },
      ".panel > .panel-body + .table, .panel > .panel-body + .table-responsive, .panel > .table + .panel-body, .panel > .table-responsive + .panel-body": {
        borderTop: "1px solid #ddd"
      },
      ".panel > .table > tbody:first-child > tr:first-child th, .panel > .table > tbody:first-child > tr:first-child td": {
        borderTop: "0"
      },
      ".panel > .table-bordered, .panel > .table-responsive > .table-bordered": {
        border: "0"
      },
      ".panel > .table-bordered > thead > tr > th:first-child, .panel > .table-responsive > .table-bordered > thead > tr > th:first-child, .panel > .table-bordered > tbody > tr > th:first-child, .panel > .table-responsive > .table-bordered > tbody > tr > th:first-child, .panel > .table-bordered > tfoot > tr > th:first-child, .panel > .table-responsive > .table-bordered > tfoot > tr > th:first-child, .panel > .table-bordered > thead > tr > td:first-child, .panel > .table-responsive > .table-bordered > thead > tr > td:first-child, .panel > .table-bordered > tbody > tr > td:first-child, .panel > .table-responsive > .table-bordered > tbody > tr > td:first-child, .panel > .table-bordered > tfoot > tr > td:first-child, .panel > .table-responsive > .table-bordered > tfoot > tr > td:first-child": {
        borderLeft: "0"
      },
      ".panel > .table-bordered > thead > tr > th:last-child, .panel > .table-responsive > .table-bordered > thead > tr > th:last-child, .panel > .table-bordered > tbody > tr > th:last-child, .panel > .table-responsive > .table-bordered > tbody > tr > th:last-child, .panel > .table-bordered > tfoot > tr > th:last-child, .panel > .table-responsive > .table-bordered > tfoot > tr > th:last-child, .panel > .table-bordered > thead > tr > td:last-child, .panel > .table-responsive > .table-bordered > thead > tr > td:last-child, .panel > .table-bordered > tbody > tr > td:last-child, .panel > .table-responsive > .table-bordered > tbody > tr > td:last-child, .panel > .table-bordered > tfoot > tr > td:last-child, .panel > .table-responsive > .table-bordered > tfoot > tr > td:last-child": {
        borderRight: "0"
      },
      ".panel > .table-bordered > thead > tr:first-child > td, .panel > .table-responsive > .table-bordered > thead > tr:first-child > td, .panel > .table-bordered > tbody > tr:first-child > td, .panel > .table-responsive > .table-bordered > tbody > tr:first-child > td, .panel > .table-bordered > thead > tr:first-child > th, .panel > .table-responsive > .table-bordered > thead > tr:first-child > th, .panel > .table-bordered > tbody > tr:first-child > th, .panel > .table-responsive > .table-bordered > tbody > tr:first-child > th": {
        borderBottom: "0"
      },
      ".panel > .table-bordered > tbody > tr:last-child > td, .panel > .table-responsive > .table-bordered > tbody > tr:last-child > td, .panel > .table-bordered > tfoot > tr:last-child > td, .panel > .table-responsive > .table-bordered > tfoot > tr:last-child > td, .panel > .table-bordered > tbody > tr:last-child > th, .panel > .table-responsive > .table-bordered > tbody > tr:last-child > th, .panel > .table-bordered > tfoot > tr:last-child > th, .panel > .table-responsive > .table-bordered > tfoot > tr:last-child > th": {
        borderBottom: "0"
      },
      ".panel > .table-responsive": {
        marginBottom: "0",
        border: "0"
      },
      ".panel-group": {
        marginBottom: 20
      },
      ".panel-group .panel": {
        marginBottom: "0",
        borderRadius: 4
      },
      ".panel-group .panel + .panel": {
        marginTop: 5
      },
      ".panel-group .panel-heading": {
        borderBottom: "0"
      },
      ".panel-group .panel-heading + .panel-collapse > .panel-body, .panel-group .panel-heading + .panel-collapse > .list-group": {
        borderTop: "1px solid #ddd"
      },
      ".panel-group .panel-footer": {
        borderTop: "0"
      },
      ".panel-group .panel-footer + .panel-collapse .panel-body": {
        borderBottom: "1px solid #ddd"
      },
      ".panel-default > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: "#ddd"
      },
      ".panel-default > .panel-heading .badge": {
        color: "#f5f5f5",
        backgroundColor: "#333"
      },
      ".panel-default > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: "#ddd"
      },
      ".panel-primary > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: theme.colors.primary.main
      },
      ".panel-primary > .panel-heading .badge": {
        color: theme.colors.primary.main,
        backgroundColor: "#fff"
      },
      ".panel-primary > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: theme.colors.primary.main
      },
      ".panel-success > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: "#d6e9c6"
      },
      ".panel-success > .panel-heading .badge": {
        color: "#dff0d8",
        backgroundColor: "#3c763d"
      },
      ".panel-success > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: "#d6e9c6"
      },
      ".panel-info > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: "#bce8f1"
      },
      ".panel-info > .panel-heading .badge": {
        color: "#d9edf7",
        backgroundColor: "#31708f"
      },
      ".panel-info > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: "#bce8f1"
      },
      ".panel-warning > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: "#faebcc"
      },
      ".panel-warning > .panel-heading .badge": {
        color: "#fcf8e3",
        backgroundColor: "#8a6d3b"
      },
      ".panel-warning > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: "#faebcc"
      },
      ".panel-danger > .panel-heading + .panel-collapse > .panel-body": {
        borderTopColor: "#ebccd1"
      },
      ".panel-danger > .panel-heading .badge": {
        color: "#f2dede",
        backgroundColor: "#a94442"
      },
      ".panel-danger > .panel-footer + .panel-collapse > .panel-body": {
        borderBottomColor: "#ebccd1"
      },
      ".embed-responsive": {
        position: "relative",
        display: "block",
        height: "0",
        padding: "0",
        overflow: "hidden"
      },
      ".embed-responsive .embed-responsive-item, .embed-responsive iframe, .embed-responsive embed, .embed-responsive object, .embed-responsive video": {
        position: "absolute",
        top: "0",
        bottom: "0",
        left: "0",
        width: "100%",
        height: "100%",
        border: "0"
      },
      ".embed-responsive-16by9": {
        paddingBottom: "56.25%"
      },
      ".embed-responsive-4by3": {
        paddingBottom: "75%"
      },
      ".well": {
        minHeight: 20,
        padding: 19,
        marginBottom: 20,
        backgroundColor: "#f5f5f5",
        border: "1px solid #e3e3e3",
        borderRadius: 4,
        W: "inset 0 1px 1px rgba(0, 0, 0, .05)",
        boxShadow: "inset 0 1px 1px rgba(0, 0, 0, .05)"
      },
      ".well blockquote": {
        borderColor: "rgba(0, 0, 0, .15)",
        fallbacks: [
          {
            borderColor: "#ddd"
          }
        ]
      },
      ".well-lg": {
        padding: 24,
        borderRadius: 6
      },
      ".well-sm": {
        padding: 9,
        borderRadius: 3
      },
      ".close": {
        float: "right",
        fontSize: 21,
        fontWeight: "bold",
        lineHeight: "1",
        color: "#000",
        textShadow: "0 1px 0 #fff",
        filter: "alpha(opacity=20)",
        opacity: ".2"
      },
      ".close:hover, .close:focus": {
        color: "#000",
        textDecoration: "none",
        cursor: "pointer",
        filter: "alpha(opacity=50)",
        opacity: ".5"
      },
      "button.close": {
        W: "none",
        padding: "0",
        cursor: "pointer",
        background: "transparent",
        border: "0"
      },
      ".modal-open": {
        "& #root": {
          filter: "blur(2px)"
        },
        overflow: "hidden",
          transition: theme.transition.time
      },
      ".modal": {
        position: "fixed",
        top: "0",
        right: "0",
        bottom: "0",
        left: "0",
        zIndex: "1050",
        display: "none",
        overflow: "hidden",
        W: "touch",
        outline: "0"
      },
      ".modal.fade .modal-dialog": {
        W: "translate(0, -25%)",
        O: "translate(0, -25%)",
        transition: "transform .3s ease-out",
        fallbacks: [
          {
            O: "-o-transform .3s ease-out"
          },
          {
            W: "-webkit-transform .3s ease-out"
          }
        ],
        M: "translate(0, -25%)",
        transform: "translate(0, -25%)"
      },
      ".modal.in .modal-dialog": {
        W: "translate(0, 0)",
        M: "translate(0, 0)",
        O: "translate(0, 0)",
        transform: "translate(0, 0)"
      },
      ".modal-open .modal": {
        overflowX: "hidden",
        overflowY: "auto"
      },
      ".modal-dialog": {
        position: "relative",
        width: "auto",
        margin: 10
      },
      ".modal-content": {
        position: "relative",
        backgroundColor: "#fff",
        W: "0 3px 9px rgba(0, 0, 0, .5)",
        backgroundClip: "padding-box",
        border: "1px solid rgba(0, 0, 0, .2)",
        fallbacks: [
          {
            W: "padding-box"
          },
          {
            border: "1px solid #999"
          }
        ],
        borderRadius: 6,
        outline: "0",
        boxShadow: "0 3px 9px rgba(0, 0, 0, .5)"
      },
      ".modal-backdrop": {
        position: "fixed",
        top: "0",
        right: "0",
        bottom: "0",
        left: "0",
        zIndex: "1040",
        backgroundColor: "#000"
      },
      ".modal-backdrop.fade": {
        filter: "alpha(opacity=0)",
        opacity: "0"
      },
      ".modal-backdrop.in": {
        filter: "alpha(opacity=50)",
        opacity: ".5"
      },
      ".modal-header": {
        padding: 15,
        borderBottom: "1px solid #e5e5e5"
      },
      ".modal-header .close": {
        marginTop: -2
      },
      ".modal-title": {
        margin: "0",
        lineHeight: "1.42857143"
      },
      ".modal-body": {
        position: "relative",
        padding: 15
      },
      ".modal-footer": {
        padding: 15,
        textAlign: "right",
        borderTop: "1px solid #e5e5e5"
      },
      ".modal-footer .btn + .btn": {
        marginBottom: "0",
        marginLeft: 5
      },
      ".modal-footer .btn-group .btn + .btn": {
        marginLeft: -1
      },
      ".modal-footer .btn-block + .btn-block": {
        marginLeft: "0"
      },
      ".modal-scrollbar-measure": {
        position: "absolute",
        top: -9999,
        width: 50,
        height: 50,
        overflow: "scroll"
      },
      ".tooltip": {
        position: "absolute",
        zIndex: "1070",
        display: "block",
        fontFamily: '"Helvetica Neue", Helvetica, Arial, sans-serif',
        fontSize: 12,
        fontStyle: "normal",
        fontWeight: "normal",
        lineHeight: "1.42857143",
        textAlign: "start",
        fallbacks: [
          {
            textAlign: "left"
          }
        ],
        textDecoration: "none",
        textShadow: "none",
        textTransform: "none",
        letterSpacing: "normal",
        wordBreak: "normal",
        wordSpacing: "normal",
        wordWrap: "normal",
        whiteSpace: "normal",
        filter: "alpha(opacity=0)",
        opacity: "0",
        lineBreak: "auto"
      },
      ".tooltip.in": {
        filter: "alpha(opacity=90)",
        opacity: ".9"
      },
      ".tooltip.top": {
        padding: "5px 0",
        marginTop: -3
      },
      ".tooltip.right": {
        padding: "0 5px",
        marginLeft: 3
      },
      ".tooltip.bottom": {
        padding: "5px 0",
        marginTop: 3
      },
      ".tooltip.left": {
        padding: "0 5px",
        marginLeft: -3
      },
      ".tooltip-inner": {
        maxWidth: 200,
        padding: "3px 8px",
        color: "#fff",
        textAlign: "center",
        backgroundColor: "#000",
        borderRadius: 4
      },
      ".tooltip-arrow": {
        position: "absolute",
        width: "0",
        height: "0",
        borderColor: "transparent",
        borderStyle: "solid"
      },
      ".tooltip.top .tooltip-arrow": {
        bottom: "0",
        left: "50%",
        marginLeft: -5,
        borderWidth: "5px 5px 0",
        borderTopColor: "#000"
      },
      ".tooltip.top-left .tooltip-arrow": {
        right: 5,
        bottom: "0",
        marginBottom: -5,
        borderWidth: "5px 5px 0",
        borderTopColor: "#000"
      },
      ".tooltip.top-right .tooltip-arrow": {
        bottom: "0",
        left: 5,
        marginBottom: -5,
        borderWidth: "5px 5px 0",
        borderTopColor: "#000"
      },
      ".tooltip.right .tooltip-arrow": {
        top: "50%",
        left: "0",
        marginTop: -5,
        borderWidth: "5px 5px 5px 0",
        borderRightColor: "#000"
      },
      ".tooltip.left .tooltip-arrow": {
        top: "50%",
        right: "0",
        marginTop: -5,
        borderWidth: "5px 0 5px 5px",
        borderLeftColor: "#000"
      },
      ".tooltip.bottom .tooltip-arrow": {
        top: "0",
        left: "50%",
        marginLeft: -5,
        borderWidth: "0 5px 5px",
        borderBottomColor: "#000"
      },
      ".tooltip.bottom-left .tooltip-arrow": {
        top: "0",
        right: 5,
        marginTop: -5,
        borderWidth: "0 5px 5px",
        borderBottomColor: "#000"
      },
      ".tooltip.bottom-right .tooltip-arrow": {
        top: "0",
        left: 5,
        marginTop: -5,
        borderWidth: "0 5px 5px",
        borderBottomColor: "#000"
      },
      ".popover": {
        position: "absolute",
        top: "0",
        left: "0",
        zIndex: "1060",
        display: "none",
        maxWidth: 276,
        padding: 1,
        fontFamily: '"Helvetica Neue", Helvetica, Arial, sans-serif',
        fontSize: 14,
        fontStyle: "normal",
        fontWeight: "normal",
        lineHeight: "1.42857143",
        textAlign: "start",
        fallbacks: [
          {
            W: "padding-box"
          },
          {
            border: "1px solid #ccc"
          },
          {
            textAlign: "left"
          }
        ],
        textDecoration: "none",
        textShadow: "none",
        textTransform: "none",
        letterSpacing: "normal",
        wordBreak: "normal",
        wordSpacing: "normal",
        wordWrap: "normal",
        whiteSpace: "normal",
        backgroundColor: "#fff",
        W: "0 5px 10px rgba(0, 0, 0, .2)",
        backgroundClip: "padding-box",
        border: "1px solid rgba(0, 0, 0, .2)",
        borderRadius: 6,
        boxShadow: "0 5px 10px rgba(0, 0, 0, .2)",
        lineBreak: "auto"
      },
      ".popover.top": {
        marginTop: -10
      },
      ".popover.right": {
        marginLeft: 10
      },
      ".popover.bottom": {
        marginTop: 10
      },
      ".popover.left": {
        marginLeft: -10
      },
      ".popover-title": {
        padding: "8px 14px",
        margin: "0",
        fontSize: 14,
        backgroundColor: "#f7f7f7",
        borderBottom: "1px solid #ebebeb",
        borderRadius: "5px 5px 0 0"
      },
      ".popover-content": {
        padding: "9px 14px"
      },
      ".popover > .arrow, .popover > .arrow:after": {
        position: "absolute",
        display: "block",
        width: "0",
        height: "0",
        borderColor: "transparent",
        borderStyle: "solid"
      },
      ".popover > .arrow": {
        borderWidth: 11
      },
      ".popover > .arrow:after": {
        content: '""',
        borderWidth: 10
      },
      ".popover.top > .arrow": {
        bottom: -11,
        left: "50%",
        marginLeft: -11,
        borderTopColor: "rgba(0, 0, 0, .25)",
        fallbacks: [
          {
            borderTopColor: "#999"
          }
        ],
        borderBottomWidth: "0"
      },
      ".popover.top > .arrow:after": {
        bottom: 1,
        marginLeft: -10,
        content: '" "',
        borderTopColor: "#fff",
        borderBottomWidth: "0"
      },
      ".popover.right > .arrow": {
        top: "50%",
        left: -11,
        marginTop: -11,
        borderRightColor: "rgba(0, 0, 0, .25)",
        fallbacks: [
          {
            borderRightColor: "#999"
          }
        ],
        borderLeftWidth: "0"
      },
      ".popover.right > .arrow:after": {
        bottom: -10,
        left: 1,
        content: '" "',
        borderRightColor: "#fff",
        borderLeftWidth: "0"
      },
      ".popover.bottom > .arrow": {
        top: -11,
        left: "50%",
        marginLeft: -11,
        borderTopWidth: "0",
        borderBottomColor: "rgba(0, 0, 0, .25)",
        fallbacks: [
          {
            borderBottomColor: "#999"
          }
        ]
      },
      ".popover.bottom > .arrow:after": {
        top: 1,
        marginLeft: -10,
        content: '" "',
        borderTopWidth: "0",
        borderBottomColor: "#fff"
      },
      ".popover.left > .arrow": {
        top: "50%",
        right: -11,
        marginTop: -11,
        borderRightWidth: "0",
        borderLeftColor: "rgba(0, 0, 0, .25)",
        fallbacks: [
          {
            borderLeftColor: "#999"
          }
        ]
      },
      ".popover.left > .arrow:after": {
        right: 1,
        bottom: -10,
        content: '" "',
        borderRightWidth: "0",
        borderLeftColor: "#fff"
      },
      ".carousel": {
        position: "relative"
      },
      ".carousel-inner": {
        position: "relative",
        width: "100%",
        overflow: "hidden"
      },
      ".carousel-inner > .item": {
        position: "relative",
        display: "none",
        W: ".6s ease-in-out left",
        O: ".6s ease-in-out left",
        transition: ".6s ease-in-out left"
      },
      ".carousel-inner > .item > img, .carousel-inner > .item > a > img": {
        lineHeight: "1"
      },
      "@media all and (transform-3d), (-webkit-transform-3d)": {
        ".carousel-inner > .item": {
          W: 1000,
          O: "-o-transform .6s ease-in-out",
          transition: "transform .6s ease-in-out",
          fallbacks: [
            {
              W: "hidden"
            },
            {
              W: "-webkit-transform .6s ease-in-out"
            }
          ],
          backfaceVisibility: "hidden",
          perspective: 1000
        },
        ".carousel-inner > .item.next, .carousel-inner > .item.active.right": {
          left: "0",
          W: "translate3d(100%, 0, 0)",
          transform: "translate3d(100%, 0, 0)"
        },
        ".carousel-inner > .item.prev, .carousel-inner > .item.active.left": {
          left: "0",
          W: "translate3d(-100%, 0, 0)",
          transform: "translate3d(-100%, 0, 0)"
        },
        ".carousel-inner > .item.next.left, .carousel-inner > .item.prev.right, .carousel-inner > .item.active": {
          left: "0",
          W: "translate3d(0, 0, 0)",
          transform: "translate3d(0, 0, 0)"
        }
      },
      ".carousel-inner > .active, .carousel-inner > .next, .carousel-inner > .prev": {
        display: "block"
      },
      ".carousel-inner > .active": {
        left: "0"
      },
      ".carousel-inner > .next, .carousel-inner > .prev": {
        position: "absolute",
        top: "0",
        width: "100%"
      },
      ".carousel-inner > .next": {
        left: "100%"
      },
      ".carousel-inner > .prev": {
        left: "-100%"
      },
      ".carousel-inner > .next.left, .carousel-inner > .prev.right": {
        left: "0"
      },
      ".carousel-inner > .active.left": {
        left: "-100%"
      },
      ".carousel-inner > .active.right": {
        left: "100%"
      },
      ".carousel-control": {
        position: "absolute",
        top: "0",
        bottom: "0",
        left: "0",
        width: "15%",
        fontSize: 20,
        color: "#fff",
        textAlign: "center",
        textShadow: "0 1px 2px rgba(0, 0, 0, .6)",
        backgroundColor: "rgba(0, 0, 0, 0)",
        filter: "alpha(opacity=50)",
        opacity: ".5"
      },
      ".carousel-control.left": {
        backgroundImage:
          "linear-gradient(to right, rgba(0, 0, 0, .5) 0%, rgba(0, 0, 0, .0001) 100%)",
        fallbacks: [
          {
            backgroundImage:
              "-webkit-gradient(linear, left top, right top, from(rgba(0, 0, 0, .5)), to(rgba(0, 0, 0, .0001)))"
          },
          {
            backgroundImage:
              "-o-linear-gradient(left, rgba(0, 0, 0, .5) 0%, rgba(0, 0, 0, .0001) 100%)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(left, rgba(0, 0, 0, .5) 0%, rgba(0, 0, 0, .0001) 100%)"
          }
        ],
        filter:
          "progid:DXImageTransform.Microsoft.gradient(startColorstr='#80000000', endColorstr='#00000000', GradientType=1)",
        backgroundRepeat: "repeat-x"
      },
      ".carousel-control.right": {
        right: "0",
        left: "auto",
        backgroundImage:
          "linear-gradient(to right, rgba(0, 0, 0, .0001) 0%, rgba(0, 0, 0, .5) 100%)",
        fallbacks: [
          {
            backgroundImage:
              "-webkit-gradient(linear, left top, right top, from(rgba(0, 0, 0, .0001)), to(rgba(0, 0, 0, .5)))"
          },
          {
            backgroundImage:
              "-o-linear-gradient(left, rgba(0, 0, 0, .0001) 0%, rgba(0, 0, 0, .5) 100%)"
          },
          {
            backgroundImage:
              "-webkit-linear-gradient(left, rgba(0, 0, 0, .0001) 0%, rgba(0, 0, 0, .5) 100%)"
          }
        ],
        filter:
          "progid:DXImageTransform.Microsoft.gradient(startColorstr='#00000000', endColorstr='#80000000', GradientType=1)",
        backgroundRepeat: "repeat-x"
      },
      ".carousel-control:hover, .carousel-control:focus": {
        color: "#fff",
        textDecoration: "none",
        filter: "alpha(opacity=90)",
        outline: "0",
        opacity: ".9"
      },
      ".carousel-control .icon-prev, .carousel-control .icon-next, .carousel-control .glyphicon-chevron-left, .carousel-control .glyphicon-chevron-right": {
        position: "absolute",
        top: "50%",
        zIndex: "5",
        display: "inline-block",
        marginTop: -10
      },
      ".carousel-control .icon-prev, .carousel-control .glyphicon-chevron-left": {
        left: "50%",
        marginLeft: -10
      },
      ".carousel-control .icon-next, .carousel-control .glyphicon-chevron-right": {
        right: "50%",
        marginRight: -10
      },
      ".carousel-control .icon-prev, .carousel-control .icon-next": {
        width: 20,
        height: 20,
        fontFamily: "serif",
        lineHeight: "1"
      },
      ".carousel-control .icon-prev:before": {
        content: "'\\2039'"
      },
      ".carousel-control .icon-next:before": {
        content: "'\\203a'"
      },
      ".carousel-indicators": {
        position: "absolute",
        bottom: 10,
        left: "50%",
        zIndex: "15",
        width: "60%",
        paddingLeft: "0",
        marginLeft: "-30%",
        textAlign: "center",
        listStyle: "none"
      },
      ".carousel-indicators li": {
        display: "inline-block",
        width: 10,
        height: 10,
        margin: 1,
        textIndent: -999,
        cursor: "pointer",
        backgroundColor: "rgba(0, 0, 0, 0)",
        fallbacks: [
          {
            backgroundColor: "#000 9"
          }
        ],
        border: "1px solid #fff",
        borderRadius: 10
      },
      ".carousel-indicators .active": {
        width: 12,
        height: 12,
        margin: "0",
        backgroundColor: "#fff"
      },
      ".carousel-caption": {
        position: "absolute",
        right: "15%",
        bottom: 20,
        left: "15%",
        zIndex: "10",
        paddingTop: 20,
        paddingBottom: 20,
        color: "#fff",
        textAlign: "center",
        textShadow: "0 1px 2px rgba(0, 0, 0, .6)"
      },
      ".carousel-caption .btn": {
        textShadow: "none"
      },
      ".clearfix:before, .clearfix:after, .dl-horizontal dd:before, .dl-horizontal dd:after, .container:before, .container:after, .container-fluid:before, .container-fluid:after, .row:before, .row:after, .form-horizontal .form-group:before, .form-horizontal .form-group:after, .btn-toolbar:before, .btn-toolbar:after, .btn-group-vertical > .btn-group:before, .btn-group-vertical > .btn-group:after, .nav:before, .nav:after, .navbar:before, .navbar:after, .navbar-header:before, .navbar-header:after, .navbar-collapse:before, .navbar-collapse:after, .pager:before, .pager:after, .panel-body:before, .panel-body:after, .modal-header:before, .modal-header:after, .modal-footer:before, .modal-footer:after": {
        display: "table",
        content: '" "'
      },
      ".clearfix:after, .dl-horizontal dd:after, .container:after, .container-fluid:after, .row:after, .form-horizontal .form-group:after, .btn-toolbar:after, .btn-group-vertical > .btn-group:after, .nav:after, .navbar:after, .navbar-header:after, .navbar-collapse:after, .pager:after, .panel-body:after, .modal-header:after, .modal-footer:after": {
        clear: "both"
      },
      ".center-block": {
        display: "block",
        marginRight: "auto",
        marginLeft: "auto"
      },
      ".pull-right": {
        float: "right !important"
      },
      ".pull-left": {
        float: "left !important"
      },
      ".hide": {
        display: "none !important"
      },
      ".show": {
        display: "block !important"
      },
      ".invisible": {
        visibility: "hidden"
      },
      ".text-hide": {
        font: "0/0 a",
        color: "transparent",
        textShadow: "none",
        backgroundColor: "transparent",
        border: "0"
      },
      ".hidden": {
        display: "none !important"
      },
      ".affix": {
        position: "fixed"
      },
      "@-ms-viewport": {
        width: "device-width"
      },
      ".visible-xs, .visible-sm, .visible-md, .visible-lg": {
        display: "none !important"
      },
      ".visible-xs-block, .visible-xs-inline, .visible-xs-inline-block, .visible-sm-block, .visible-sm-inline, .visible-sm-inline-block, .visible-md-block, .visible-md-inline, .visible-md-inline-block, .visible-lg-block, .visible-lg-inline, .visible-lg-inline-block": {
        display: "none !important"
      },
      "@media (min-width: 768px) and (max-width: 991px)": {
        ".visible-sm": {
          display: "block !important"
        },
        "table.visible-sm": {
          display: "table !important"
        },
        "tr.visible-sm": {
          display: "table-row !important"
        },
        "th.visible-sm, td.visible-sm": {
          display: "table-cell !important"
        },
        ".visible-sm-block": {
          display: "block !important"
        },
        ".visible-sm-inline": {
          display: "inline !important"
        },
        ".visible-sm-inline-block": {
          display: "inline-block !important"
        },
        ".hidden-sm": {
          display: "none !important"
        }
      },
      "@media (min-width: 992px) and (max-width: 1199px)": {
        ".visible-md": {
          display: "block !important"
        },
        "table.visible-md": {
          display: "table !important"
        },
        "tr.visible-md": {
          display: "table-row !important"
        },
        "th.visible-md, td.visible-md": {
          display: "table-cell !important"
        },
        ".visible-md-block": {
          display: "block !important"
        },
        ".visible-md-inline": {
          display: "inline !important"
        },
        ".visible-md-inline-block": {
          display: "inline-block !important"
        },
        ".hidden-md": {
          display: "none !important"
        }
      },
      ".visible-print": {
        display: "none !important"
      },
      ".visible-print-block": {
        display: "none !important"
      },
      ".visible-print-inline": {
        display: "none !important"
      },
      ".visible-print-inline-block": {
        display: "none !important"
      },
      ".draftJsToolbar__buttonWrapper__1Dmqh, .draftJsEmojiPlugin__buttonWrapper__1Dmqh": {
        display: "inline-block" as "inline-block"
      },
      ".draftJsToolbar__button__qi1gf, .draftJsEmojiPlugin__button__qi1gf": {
        background: "transparent",
        color: "#888",
        fontSize: 18,
        border: "0",
        paddingTop: 5,
        verticalAlign: "bottom" as "bottom",
        height: 34,
        width: 36
      },
      ".draftJsToolbar__button__qi1gf svg, .draftJsEmojiPlugin__button__qi1gf svg": {
        fill: "#888"
      },
      ".draftJsToolbar__button__qi1gf:hover, .draftJsToolbar__button__qi1gf:focus, .draftJsEmojiPlugin__button__qi1gf:hover, .draftJsEmojiPlugin__button__qi1gf:focus": {
        background: "#f3f3f3",
        outline: "0"
      },
      ".draftJsToolbar__active__3qcpF": {
        background: "#efefef",
        color: "#444"
      },
      ".draftJsToolbar__active__3qcpF svg": {
        fill: "#444"
      },
      ".draftJsToolbar__separator__3U7qt": {
        display: "inline-block",
        borderRight: "1px solid #ddd",
        height: 24,
        margin: "0 0.5em"
      },
      ".draftJsToolbar__toolbar__dNtBH": {
        left: "50%",
        W: "translate(-50%) scale(0)",
        transform: "translate(-50%) scale(0)",
        position: "absolute",
        border: "1px solid #ddd",
        background: "#fff",
        borderRadius: 2,
        boxShadow: "0px 1px 3px 0px rgba(220,220,220,1)",
        zIndex: "2",
        boxSizing: "border-box"
      },
      ".draftJsToolbar__toolbar__dNtBH:after, .draftJsToolbar__toolbar__dNtBH:before": {
        top: "100%",
        left: "50%",
        border: "solid transparent",
        content: '" "',
        height: "0",
        width: "0",
        position: "absolute",
        pointerEvents: "none"
      },
      ".draftJsToolbar__toolbar__dNtBH:after, .draftJsEmojiPlugin__alignmentTool__2mkQr:after": {
        borderColor: "rgba(255, 255, 255, 0)",
        borderTopColor: "#fff",
        borderWidth: 4,
        marginLeft: -4
      },
      ".draftJsToolbar__toolbar__dNtBH:before, .draftJsEmojiPlugin__alignmentTool__2mkQr:before": {
        borderColor: "rgba(221, 221, 221, 0)",
        borderTopColor: "#ddd",
        borderWidth: 6,
        marginLeft: -6
      },
      ".draftJsEmojiPlugin__alignmentTool__2mkQr": {
        background: "#fff",
        border: "1px solid #ddd",
        borderRadius: 2,
        boxShadow: "0px 1px 3px 0px rgba(220,220,220,1)",
        display: "flex" as "flex",
        position: "absolute" as "absolute"
      }
    }
  } as any);
