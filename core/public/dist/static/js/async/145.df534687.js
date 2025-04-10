"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["145"],{3218:function(e,t,a){let i;a.r(t),a.d(t,{default:()=>en});var l=a(9454),r=a(2403),o=a(6419),n=a(6988),d=a(209);let s=(0,d.aZ)({__name:"index",props:(0,d.Vf)({width:{default:240},prefix:{type:Boolean,default:!1}},{value:{default:""},valueModifiers:{}}),emits:(0,d.Vf)(["search"],["update:value"]),setup(e,t){let{expose:a,emit:i}=t,l=(0,d.iH)(null),r=(0,d.tT)(e,"value"),o=()=>{i("search",r.value)};return a({focus:()=>{var e;null===(e=l.value)||void 0===e||e.focus()}}),(e,t)=>{let a=n.Z;return(0,d.wg)(),(0,d.j4)(a,(0,d.dG)({ref_key:"inputRef",ref:l,value:r.value,"onUpdate:value":t[0]||(t[0]=e=>r.value=e)},e.$attrs,{class:"bt-search",style:{width:`${e.width}px`},onKeyup:(0,d.D2)(o,["enter"])}),(0,d.Nv)({_:2},[e.prefix?{name:"prefix",fn:(0,d.w5)(()=>[(0,d._)("div",{class:"flex items-center cursor-pointer",onClick:o},t[1]||(t[1]=[(0,d._)("i",{class:"i-mdi-search text-16px"},null,-1)]))]),key:"0"}:void 0,e.prefix?void 0:{name:"suffix",fn:(0,d.w5)(()=>[(0,d._)("div",{class:"flex items-center cursor-pointer",onClick:o},t[2]||(t[2]=[(0,d._)("i",{class:"i-mdi-search text-16px"},null,-1)]))]),key:"1"}]),1040,["value","style"])}}});var u=a(3447),c=a(5083),p=a(9226),h=a(6154),m=a(4131),b=a(1321),v=a(4124),f=a(9241),w=a(6169),g=a(1844),x=a(2249),_=a(8282),y=a(363),k=a(8755);let S={buttonHeightSmall:"14px",buttonHeightMedium:"18px",buttonHeightLarge:"22px",buttonWidthSmall:"14px",buttonWidthMedium:"18px",buttonWidthLarge:"22px",buttonWidthPressedSmall:"20px",buttonWidthPressedMedium:"24px",buttonWidthPressedLarge:"28px",railHeightSmall:"18px",railHeightMedium:"22px",railHeightLarge:"26px",railWidthSmall:"32px",railWidthMedium:"40px",railWidthLarge:"48px"},U={name:"Switch",common:k.Z,self:function(e){let{primaryColor:t,opacityDisabled:a,borderRadius:i,textColor3:l}=e;return Object.assign(Object.assign({},S),{iconColor:l,textColor:"white",loadingColor:t,opacityDisabled:a,railColor:"rgba(0, 0, 0, .14)",railColorActive:t,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:i,railBorderRadiusMedium:i,railBorderRadiusLarge:i,buttonBorderRadiusSmall:i,buttonBorderRadiusMedium:i,buttonBorderRadiusLarge:i,boxShadowFocus:`0 0 0 2px ${(0,y.zX)(t,{alpha:.2})}`})}};var B=a(8758);let W=(0,x.cB)("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[(0,x.cE)("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),(0,x.cE)("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),(0,x.cE)("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),(0,x.cB)("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[(0,B.c)({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),(0,x.cE)("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),(0,x.cE)("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),(0,x.cE)("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),(0,x.c)("&:focus",[(0,x.cE)("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),(0,x.cM)("round",[(0,x.cE)("rail","border-radius: calc(var(--n-rail-height) / 2);",[(0,x.cE)("button","border-radius: calc(var(--n-button-height) / 2);")])]),(0,x.u4)("disabled",[(0,x.u4)("icon",[(0,x.cM)("rubber-band",[(0,x.cM)("pressed",[(0,x.cE)("rail",[(0,x.cE)("button","max-width: var(--n-button-width-pressed);")])]),(0,x.cE)("rail",[(0,x.c)("&:active",[(0,x.cE)("button","max-width: var(--n-button-width-pressed);")])]),(0,x.cM)("active",[(0,x.cM)("pressed",[(0,x.cE)("rail",[(0,x.cE)("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),(0,x.cE)("rail",[(0,x.c)("&:active",[(0,x.cE)("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),(0,x.cM)("active",[(0,x.cE)("rail",[(0,x.cE)("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),(0,x.cE)("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[(0,x.cE)("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[(0,B.c)()]),(0,x.cE)("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),(0,x.cM)("active",[(0,x.cE)("rail","background-color: var(--n-rail-color-active);")]),(0,x.cM)("loading",[(0,x.cE)("rail",`
 cursor: wait;
 `)]),(0,x.cM)("disabled",[(0,x.cE)("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),E=Object.assign(Object.assign({},b.Z.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]}),$=(0,d.aZ)({name:"Switch",props:E,slots:Object,setup(e){void 0===i&&(i="undefined"==typeof CSS||void 0!==CSS.supports&&CSS.supports("width","max(1px)"));let{mergedClsPrefixRef:t,inlineThemeDisabled:a}=(0,v.ZP)(e),l=(0,b.Z)("Switch","-switch",W,U,e,t),r=(0,f.Z)(e),{mergedSizeRef:o,mergedDisabledRef:n}=r,s=(0,d.iH)(e.defaultValue),u=(0,d.Vh)(e,"value"),h=(0,p.Z)(u,s),m=(0,d.Fl)(()=>h.value===e.checkedValue),_=(0,d.iH)(!1),y=(0,d.iH)(!1),k=(0,d.Fl)(()=>{let{railStyle:t}=e;if(t)return t({focused:y.value,checked:m.value})});function S(t){let{"onUpdate:value":a,onChange:i,onUpdateValue:l}=e,{nTriggerFormInput:o,nTriggerFormChange:n}=r;a&&(0,g.R)(a,t),l&&(0,g.R)(l,t),i&&(0,g.R)(i,t),s.value=t,o(),n()}let B=(0,d.Fl)(()=>{let e,t,a;let{value:r}=o,{self:{opacityDisabled:n,railColor:d,railColorActive:s,buttonBoxShadow:u,buttonColor:p,boxShadowFocus:h,loadingColor:m,textColor:b,iconColor:v,[(0,x.Tl)("buttonHeight",r)]:f,[(0,x.Tl)("buttonWidth",r)]:w,[(0,x.Tl)("buttonWidthPressed",r)]:g,[(0,x.Tl)("railHeight",r)]:_,[(0,x.Tl)("railWidth",r)]:y,[(0,x.Tl)("railBorderRadius",r)]:k,[(0,x.Tl)("buttonBorderRadius",r)]:S},common:{cubicBezierEaseInOut:U}}=l.value;return i?(e=`calc((${_} - ${f}) / 2)`,t=`max(${_}, ${f})`,a=`max(${y}, calc(${y} + ${f} - ${_}))`):(e=(0,c.BL)(((0,c.fQ)(_)-(0,c.fQ)(f))/2),t=(0,c.BL)(Math.max((0,c.fQ)(_),(0,c.fQ)(f))),a=(0,c.fQ)(_)>(0,c.fQ)(f)?y:(0,c.BL)((0,c.fQ)(y)+(0,c.fQ)(f)-(0,c.fQ)(_))),{"--n-bezier":U,"--n-button-border-radius":S,"--n-button-box-shadow":u,"--n-button-color":p,"--n-button-width":w,"--n-button-width-pressed":g,"--n-button-height":f,"--n-height":t,"--n-offset":e,"--n-opacity-disabled":n,"--n-rail-border-radius":k,"--n-rail-color":d,"--n-rail-color-active":s,"--n-rail-height":_,"--n-rail-width":y,"--n-width":a,"--n-box-shadow-focus":h,"--n-loading-color":m,"--n-text-color":b,"--n-icon-color":v}}),E=a?(0,w.F)("switch",(0,d.Fl)(()=>o.value[0]),B,e):void 0;return{handleClick:function(){e.loading||n.value||(h.value!==e.checkedValue?S(e.checkedValue):S(e.uncheckedValue))},handleBlur:function(){y.value=!1,function(){let{nTriggerFormBlur:e}=r;e()}(),_.value=!1},handleFocus:function(){y.value=!0,function(){let{nTriggerFormFocus:e}=r;e()}()},handleKeyup:function(t){e.loading||n.value||" "!==t.key||(h.value!==e.checkedValue?S(e.checkedValue):S(e.uncheckedValue),_.value=!1)},handleKeydown:function(t){e.loading||n.value||" "!==t.key||(t.preventDefault(),_.value=!0)},mergedRailStyle:k,pressed:_,mergedClsPrefix:t,mergedValue:h,checked:m,mergedDisabled:n,cssVars:a?void 0:B,themeClass:null==E?void 0:E.themeClass,onRender:null==E?void 0:E.onRender}},render(){let{mergedClsPrefix:e,mergedDisabled:t,checked:a,mergedRailStyle:i,onRender:l,$slots:r}=this;null==l||l();let{checked:o,unchecked:n,icon:s,"checked-icon":u,"unchecked-icon":c}=r,p=!((0,_.aD)(s)&&(0,_.aD)(u)&&(0,_.aD)(c));return(0,d.h)("div",{role:"switch","aria-checked":a,class:[`${e}-switch`,this.themeClass,p&&`${e}-switch--icon`,a&&`${e}-switch--active`,t&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},(0,d.h)("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:i},(0,_.K9)(o,t=>(0,_.K9)(n,a=>t||a?(0,d.h)("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},(0,d.h)("div",{class:`${e}-switch__rail-placeholder`},(0,d.h)("div",{class:`${e}-switch__button-placeholder`}),t),(0,d.h)("div",{class:`${e}-switch__rail-placeholder`},(0,d.h)("div",{class:`${e}-switch__button-placeholder`}),a)):null)),(0,d.h)("div",{class:`${e}-switch__button`},(0,_.K9)(s,t=>(0,_.K9)(u,a=>(0,_.K9)(c,i=>(0,d.h)(h.Z,null,{default:()=>this.loading?(0,d.h)(m.Z,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(a||t)?(0,d.h)("div",{class:`${e}-switch__button-icon`,key:a?"checked-icon":"icon"},a||t):!this.checked&&(i||t)?(0,d.h)("div",{class:`${e}-switch__button-icon`,key:i?"unchecked-icon":"icon"},i||t):null})))),(0,_.K9)(o,t=>t&&(0,d.h)("div",{key:"checked",class:`${e}-switch__checked`},t)),(0,_.K9)(n,t=>t&&(0,d.h)("div",{key:"unchecked",class:`${e}-switch__unchecked`},t)))))}});var Z=a(9506),C=a(5702),z=a(3312),M=a(6949),F=a(4203);let R=e=>F.e.get("/mailbox/list",{params:e}),V=e=>F.e.post("/mailbox/create",e,{fetchOptions:{successMessage:!0}}),A=e=>F.e.post("/mailbox/update",e,{fetchOptions:{successMessage:!0}}),H=e=>F.e.post("/mailbox/delete",e,{fetchOptions:{successMessage:!0}});var P=a(5476);let K={class:"flex"},q={class:"w-0 flex-1 flex items-center"},T={class:"min-w-0"},j=(0,d.aZ)({__name:"index",props:{value:{default:""}},setup(e){let t=(0,d.iH)(!1),a=(0,d.Fl)(()=>t.value?e.value:"**********"),i=()=>{t.value=!t.value},l=async()=>{(0,C.zp)(`${e.value}`)};return(e,r)=>{let o=P.ZP;return(0,d.wg)(),(0,d.iD)("div",K,[(0,d._)("div",q,[(0,d._)("div",T,[(0,d.Wm)(o,{class:"max-w-full",tooltip:{contentStyle:{maxWidth:"200px"}}},{default:(0,d.w5)(()=>[(0,d.Uk)((0,d.zw)((0,d.SU)(a)),1)]),_:1})]),(0,d._)("div",{class:"ml-6px text-15px cursor-pointer",title:"Show",onClick:i},[(0,d._)("i",{class:(0,d.C_)((0,d.SU)(t)?"i-mdi-eye-outline":"i-mdi-eye-off-outline")},null,2)]),(0,d._)("div",{class:"reset ml-6px text-15px cursor-pointer text-#333",title:"Copy",onClick:l},r[0]||(r[0]=[(0,d._)("i",{class:"i-mdi-content-copy"},null,-1)]))])])}}});var L=a(5891),D=a(55);let O=(0,d.aZ)({__name:"DomainSelect",props:(0,d.Vf)({isAll:{type:Boolean,default:!0}},{value:{default:()=>null},valueModifiers:{}}),emits:["update:value"],setup(e){let t=(0,d.tT)(e,"value"),a=(0,d.iH)(!1),i=(0,d.iH)([]);return(async()=>{try{a.value=!0;let l=await (0,D.i7)({page:1,page_size:1e4,keyword:""});(0,C.Kn)(l)&&(i.value=l.list.map(e=>({label:e.domain,value:e.domain})),e.isAll?(i.value.unshift({label:"All",value:""}),t.value=""):l.list.length>0&&!t.value&&(t.value=l.list[0].domain))}finally{a.value=!1}})(),(e,l)=>{let r=L.Z;return(0,d.wg)(),(0,d.j4)(r,{value:t.value,"onUpdate:value":l[0]||(l[0]=e=>t.value=e),loading:(0,d.SU)(a),filterable:!0,options:(0,d.SU)(i),placeholder:"请选择域名"},null,8,["value","loading","options"])}}}),Q=(0,a(4222).default)(O,[["__scopeId","data-v-7274436c"]]);var G=a(5523),N=a(4603),X=a(6883),I=a(9167);let Y=(0,x.cB)("input-group",`
 display: inline-flex;
 width: 100%;
 flex-wrap: nowrap;
 vertical-align: bottom;
`,[(0,x.c)(">",[(0,x.cB)("input",[(0,x.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,x.c)("&:not(:first-child)",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 margin-left: -1px!important;
 `)]),(0,x.cB)("button",[(0,x.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `,[(0,x.cE)("state-border, border",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `)]),(0,x.c)("&:not(:first-child)",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `,[(0,x.cE)("state-border, border",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `)])]),(0,x.c)("*",[(0,x.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `,[(0,x.c)(">",[(0,x.cB)("input",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,x.cB)("base-selection",[(0,x.cB)("base-selection-label",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,x.cB)("base-selection-tags",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,x.cE)("box-shadow, border, state-border",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `)])])]),(0,x.c)("&:not(:first-child)",`
 margin-left: -1px!important;
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `,[(0,x.c)(">",[(0,x.cB)("input",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,x.cB)("base-selection",[(0,x.cB)("base-selection-label",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,x.cB)("base-selection-tags",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,x.cE)("box-shadow, border, state-border",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `)])])])])])]),J=(0,d.aZ)({name:"InputGroup",props:{},setup(e){let{mergedClsPrefixRef:t}=(0,v.ZP)(e);return(0,I.Z)("-input-group",Y,t),{mergedClsPrefix:t}},render(){let{mergedClsPrefix:e}=this;return(0,d.h)("div",{class:`${e}-input-group`},this.$slots)}}),ee={class:"w-280px"},et={class:"w-280px"},ea={class:"w-280px"},ei={class:"w-280px"},el=(0,d.aZ)({__name:"MailboxForm",setup(e){let t=(0,d.iH)(!1),a=(0,d.Fl)(()=>t.value?"编辑邮箱":"添加邮箱"),i=(0,d.AE)("formRef"),l=(0,d.qj)({quota:5,unit:"GB",isAdmin:0,full_name:"",domain:null,password:"",active:1}),r=[{label:"GB",value:"GB"},{label:"MB",value:"MB"}],o=[{label:"普通用户",value:0},{label:"管理员",value:1}],s={full_name:{trigger:"blur",validator:()=>""!==l.full_name.trim()&&!!l.domain||Error("请输入用户名")},password:{trigger:"blur",validator:()=>{if(t.value){if(l.password&&l.password.trim().length<8)return Error("密码长度少于8位，请重新输入");if(l.password&&!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/.test(l.password))return Error("密码必须包含至少一个大写字母、一个小写字母和数字，请重新输入")}else{if(l.password.trim().length<8)return Error("密码长度少于8位，请重新输入");if(!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/.test(l.password))return Error("密码必须包含至少一个大写字母、一个小写字母和数字，请重新输入")}return!0}}},u=(e,t)=>{switch(t){case"GB":e*=1024;case"MB":e*=1048576}return e},c=()=>({full_name:l.full_name,domain:l.domain||"",password:l.password,quota:u(l.quota,l.unit),isAdmin:l.isAdmin,active:l.active}),[p,h]=(0,z.d)({onChangeState:e=>{if(e){let e=h.getState(),{row:a}=e;if(t.value=e.isEdit,a){l.full_name=a.full_name,l.domain=a.domain,l.isAdmin=a.is_admin,l.active=a.active;let[e,t]=(0,C.UO)(a.quota).split(" ");l.quota=(0,C.Dx)(e),l.unit=t}}else l.full_name="",l.domain=null,l.password="",l.quota=5,l.unit="GB",l.isAdmin=0,l.active=1},onConfirm:async()=>{var e;await (null===(e=i.value)||void 0===e?void 0:e.validate());let a=c();t.value?await A(a):await V(a),h.getState().refresh()}});return(e,u)=>{let c=n.Z,h=X.ZP,m=N.Z,b=L.Z,v=G.Z;return(0,d.wg)(),(0,d.j4)((0,d.SU)(p),{title:(0,d.SU)(a),width:"520"},{default:(0,d.w5)(()=>[(0,d.Wm)(v,{ref_key:"formRef",ref:i,class:"pt-20px",model:(0,d.SU)(l),rules:s},{default:(0,d.w5)(()=>[(0,d.Wm)(h,{label:"邮箱地址",path:"full_name"},{default:(0,d.w5)(()=>[(0,d._)("div",ee,[(0,d.Wm)(J,null,{default:(0,d.w5)(()=>[(0,d.Wm)(c,{value:(0,d.SU)(l).full_name,"onUpdate:value":u[0]||(u[0]=e=>(0,d.SU)(l).full_name=e),class:"flex-1",disabled:(0,d.SU)(t)},null,8,["value","disabled"]),(0,d.Wm)(Q,{value:(0,d.SU)(l).domain,"onUpdate:value":u[1]||(u[1]=e=>(0,d.SU)(l).domain=e),class:"flex-1","is-all":!0,disabled:(0,d.SU)(t)},null,8,["value","disabled"])]),_:1})])]),_:1}),(0,d.Wm)(h,{label:"密码",path:"password"},{default:(0,d.w5)(()=>[(0,d._)("div",et,[(0,d.Wm)(c,{value:(0,d.SU)(l).password,"onUpdate:value":u[2]||(u[2]=e=>(0,d.SU)(l).password=e),placeholder:"请输入您的邮箱密码"},null,8,["value"])])]),_:1}),(0,d.Wm)(h,{label:"配额",path:"quota"},{default:(0,d.w5)(()=>[(0,d._)("div",ea,[(0,d.Wm)(J,null,{default:(0,d.w5)(()=>[(0,d.Wm)(m,{value:(0,d.SU)(l).quota,"onUpdate:value":u[3]||(u[3]=e=>(0,d.SU)(l).quota=e),class:"flex-1",min:1,"show-button":!1},null,8,["value"]),(0,d.Wm)(b,{value:(0,d.SU)(l).unit,"onUpdate:value":u[4]||(u[4]=e=>(0,d.SU)(l).unit=e),class:"flex-1",options:r},null,8,["value"])]),_:1})])]),_:1}),(0,d.Wm)(h,{label:"用户类型",path:"userType"},{default:(0,d.w5)(()=>[(0,d._)("div",ei,[(0,d.Wm)(b,{value:(0,d.SU)(l).isAdmin,"onUpdate:value":u[5]||(u[5]=e=>(0,d.SU)(l).isAdmin=e),options:o},null,8,["value"])])]),_:1}),(0,d.Wm)(h,{label:"状态"},{default:(0,d.w5)(()=>[(0,d.Wm)($,{value:(0,d.SU)(l).active,"onUpdate:value":u[6]||(u[6]=e=>(0,d.SU)(l).active=e),"checked-value":1,"unchecked-value":0},null,8,["value"])]),_:1})]),_:1},8,["model"])]),_:1},8,["title"])}}}),er={class:"p-24px"},eo={class:"w-220px"},en=(0,d.aZ)({__name:"index",setup(e){let{tableParams:t,tableList:a,loading:i,tableTotal:n,getTableData:c}=(0,M.Z)({loading:!0,immediate:!1,params:{page:1,page_size:10,domain:null,keyword:""},fetchFn:R}),p=(0,d.iH)([{key:"username",title:"Username",minWidth:120,ellipsis:{tooltip:!0}},{key:"password",title:"Password",width:"18%",minWidth:120,render:e=>(0,d.Wm)(j,{value:e.password||"--"},null)},{key:"login",title:"Login info",render:e=>(0,d.Wm)(u.ZP,{text:!0,type:"primary",onClick:()=>{(0,C.zp)(`WebMail: ${window.location.origin}/roundcube
Username: ${e.username}
Password: ${e.password}`)}},{default:()=>[(0,d.Uk)("Copy")]})},{key:"quota",title:"Quota",width:"14%",minWidth:120,render:e=>(0,C.UO)(e.quota)},{key:"is_admin",title:"Type",width:"14%",minWidth:100,render:e=>1===e.is_admin?"Admin":"General user"},{key:"status",title:"Status",width:"12%",minWidth:80,render:e=>(0,d.Wm)($,{value:e.active,"checked-value":1,"unchecked-value":0,size:"small",onUpdateValue:t=>{v(e,t)}},null)},{title:"操作",key:"actions",align:"right",width:120,render:e=>(0,d.Wm)(Z.Z,{inline:!0},{default:()=>[(0,d.Wm)(u.ZP,{type:"primary",text:!0,onClick:()=>{f(e)}},{default:()=>[(0,d.Uk)("编辑")]}),(0,d.Wm)(u.ZP,{type:"error",text:!0,onClick:()=>{w(e)}},{default:()=>[(0,d.Uk)("删除")]})]})}]),[h,m]=(0,z.d)({component:el,state:{isEdit:!1,refresh:c}}),b=()=>{m.setState({isEdit:!1,row:null}),m.open()},v=async(e,t)=>{await A({full_name:e.full_name,domain:e.domain,password:e.password,quota:e.quota,isAdmin:e.is_admin,active:t}),e.active=t},f=e=>{m.setState({isEdit:!0,row:e}),m.open()},w=e=>{(0,C.iG)({title:"删除邮箱",content:`确定要该删除邮箱【${e.username}】吗？`,confirmText:"删除",confirmType:"error",onConfirm:async()=>{await H({email:e.username}),c()}})};return(e,m)=>{let v=o.Z,f=r.Z,w=l.Z;return(0,d.wg)(),(0,d.iD)("div",er,[m[8]||(m[8]=(0,d._)("div",{class:"mb-20px text-24px font-bold"},"MailBoxes",-1)),(0,d.Wm)(w,null,{toolsLeft:(0,d.w5)(()=>[(0,d.Wm)((0,d.SU)(u.ZP),{type:"primary",onClick:b},{default:(0,d.w5)(()=>m[7]||(m[7]=[(0,d.Uk)("添加邮箱")])),_:1})]),toolsRight:(0,d.w5)(()=>[(0,d._)("div",eo,[(0,d.Wm)(Q,{value:(0,d.SU)(t).domain,"onUpdate:value":[m[0]||(m[0]=e=>(0,d.SU)(t).domain=e),m[1]||(m[1]=()=>(0,d.SU)(c)(!0))]},null,8,["value"])]),(0,d.Wm)(s,{value:(0,d.SU)(t).keyword,"onUpdate:value":m[2]||(m[2]=e=>(0,d.SU)(t).keyword=e),width:"280",placeholder:"搜索用户名",onSearch:m[3]||(m[3]=()=>(0,d.SU)(c)(!0))},null,8,["value"])]),table:(0,d.w5)(()=>[(0,d.Wm)(v,{loading:(0,d.SU)(i),columns:(0,d.SU)(p),data:(0,d.SU)(a)},null,8,["loading","columns","data"])]),pageRight:(0,d.w5)(()=>[(0,d.Wm)(f,{page:(0,d.SU)(t).page,"onUpdate:page":m[4]||(m[4]=e=>(0,d.SU)(t).page=e),"page-size":(0,d.SU)(t).page_size,"onUpdate:pageSize":m[5]||(m[5]=e=>(0,d.SU)(t).page_size=e),"item-count":(0,d.SU)(n),onRefresh:m[6]||(m[6]=()=>(0,d.SU)(c)(!0))},null,8,["page","page-size","item-count"])]),modal:(0,d.w5)(()=>[(0,d.Wm)((0,d.SU)(h))]),_:1})])}}})}}]);