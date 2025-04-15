"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["349"],{7279:function(e,t,a){let i;a.r(t),a.d(t,{default:()=>ep});var l=a(9454),o=a(2403),r=a(6419),n=a(6988),d=a(209);let s=(0,d.aZ)({__name:"index",props:(0,d.Vf)({width:{default:240},prefix:{type:Boolean,default:!1}},{value:{default:""},valueModifiers:{}}),emits:(0,d.Vf)(["search"],["update:value"]),setup(e,t){let{expose:a,emit:i}=t,l=(0,d.iH)(null),o=(0,d.tT)(e,"value"),r=()=>{i("search",o.value)};return a({focus:()=>{var e;null===(e=l.value)||void 0===e||e.focus()}}),(e,t)=>{let a=n.Z;return(0,d.wg)(),(0,d.j4)(a,(0,d.dG)({ref_key:"inputRef",ref:l,value:o.value,"onUpdate:value":t[0]||(t[0]=e=>o.value=e)},e.$attrs,{class:"bt-search",style:{width:`${e.width}px`},onKeyup:(0,d.D2)(r,["enter"])}),(0,d.Nv)({_:2},[e.prefix?{name:"prefix",fn:(0,d.w5)(()=>[(0,d._)("div",{class:"flex items-center cursor-pointer",onClick:r},t[1]||(t[1]=[(0,d._)("i",{class:"i-mdi-search text-16px"},null,-1)]))]),key:"0"}:void 0,e.prefix?void 0:{name:"suffix",fn:(0,d.w5)(()=>[(0,d._)("div",{class:"flex items-center cursor-pointer",onClick:r},t[2]||(t[2]=[(0,d._)("i",{class:"i-mdi-search text-16px"},null,-1)]))]),key:"1"}]),1040,["value","style"])}}});var u=a(3447),c=a(5083),p=a(9226),m=a(6154),h=a(4131),b=a(1321),v=a(4124),f=a(9241),g=a(6169),w=a(1844),x=a(2249),y=a(8282),S=a(363),_=a(8755);let k={buttonHeightSmall:"14px",buttonHeightMedium:"18px",buttonHeightLarge:"22px",buttonWidthSmall:"14px",buttonWidthMedium:"18px",buttonWidthLarge:"22px",buttonWidthPressedSmall:"20px",buttonWidthPressedMedium:"24px",buttonWidthPressedLarge:"28px",railHeightSmall:"18px",railHeightMedium:"22px",railHeightLarge:"26px",railWidthSmall:"32px",railWidthMedium:"40px",railWidthLarge:"48px"},U={name:"Switch",common:_.Z,self:function(e){let{primaryColor:t,opacityDisabled:a,borderRadius:i,textColor3:l}=e;return Object.assign(Object.assign({},k),{iconColor:l,textColor:"white",loadingColor:t,opacityDisabled:a,railColor:"rgba(0, 0, 0, .14)",railColorActive:t,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:i,railBorderRadiusMedium:i,railBorderRadiusLarge:i,buttonBorderRadiusSmall:i,buttonBorderRadiusMedium:i,buttonBorderRadiusLarge:i,boxShadowFocus:`0 0 0 2px ${(0,S.zX)(t,{alpha:.2})}`})}};var B=a(8758);let W=(0,x.cB)("switch",`
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
 `)])]),E=Object.assign(Object.assign({},b.Z.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]}),Z=(0,d.aZ)({name:"Switch",props:E,slots:Object,setup(e){void 0===i&&(i="undefined"==typeof CSS||void 0!==CSS.supports&&CSS.supports("width","max(1px)"));let{mergedClsPrefixRef:t,inlineThemeDisabled:a}=(0,v.ZP)(e),l=(0,b.Z)("Switch","-switch",W,U,e,t),o=(0,f.Z)(e),{mergedSizeRef:r,mergedDisabledRef:n}=o,s=(0,d.iH)(e.defaultValue),u=(0,d.Vh)(e,"value"),m=(0,p.Z)(u,s),h=(0,d.Fl)(()=>m.value===e.checkedValue),y=(0,d.iH)(!1),S=(0,d.iH)(!1),_=(0,d.Fl)(()=>{let{railStyle:t}=e;if(t)return t({focused:S.value,checked:h.value})});function k(t){let{"onUpdate:value":a,onChange:i,onUpdateValue:l}=e,{nTriggerFormInput:r,nTriggerFormChange:n}=o;a&&(0,w.R)(a,t),l&&(0,w.R)(l,t),i&&(0,w.R)(i,t),s.value=t,r(),n()}let B=(0,d.Fl)(()=>{let e,t,a;let{value:o}=r,{self:{opacityDisabled:n,railColor:d,railColorActive:s,buttonBoxShadow:u,buttonColor:p,boxShadowFocus:m,loadingColor:h,textColor:b,iconColor:v,[(0,x.Tl)("buttonHeight",o)]:f,[(0,x.Tl)("buttonWidth",o)]:g,[(0,x.Tl)("buttonWidthPressed",o)]:w,[(0,x.Tl)("railHeight",o)]:y,[(0,x.Tl)("railWidth",o)]:S,[(0,x.Tl)("railBorderRadius",o)]:_,[(0,x.Tl)("buttonBorderRadius",o)]:k},common:{cubicBezierEaseInOut:U}}=l.value;return i?(e=`calc((${y} - ${f}) / 2)`,t=`max(${y}, ${f})`,a=`max(${S}, calc(${S} + ${f} - ${y}))`):(e=(0,c.BL)(((0,c.fQ)(y)-(0,c.fQ)(f))/2),t=(0,c.BL)(Math.max((0,c.fQ)(y),(0,c.fQ)(f))),a=(0,c.fQ)(y)>(0,c.fQ)(f)?S:(0,c.BL)((0,c.fQ)(S)+(0,c.fQ)(f)-(0,c.fQ)(y))),{"--n-bezier":U,"--n-button-border-radius":k,"--n-button-box-shadow":u,"--n-button-color":p,"--n-button-width":g,"--n-button-width-pressed":w,"--n-button-height":f,"--n-height":t,"--n-offset":e,"--n-opacity-disabled":n,"--n-rail-border-radius":_,"--n-rail-color":d,"--n-rail-color-active":s,"--n-rail-height":y,"--n-rail-width":S,"--n-width":a,"--n-box-shadow-focus":m,"--n-loading-color":h,"--n-text-color":b,"--n-icon-color":v}}),E=a?(0,g.F)("switch",(0,d.Fl)(()=>r.value[0]),B,e):void 0;return{handleClick:function(){e.loading||n.value||(m.value!==e.checkedValue?k(e.checkedValue):k(e.uncheckedValue))},handleBlur:function(){S.value=!1,function(){let{nTriggerFormBlur:e}=o;e()}(),y.value=!1},handleFocus:function(){S.value=!0,function(){let{nTriggerFormFocus:e}=o;e()}()},handleKeyup:function(t){e.loading||n.value||" "!==t.key||(m.value!==e.checkedValue?k(e.checkedValue):k(e.uncheckedValue),y.value=!1)},handleKeydown:function(t){e.loading||n.value||" "!==t.key||(t.preventDefault(),y.value=!0)},mergedRailStyle:_,pressed:y,mergedClsPrefix:t,mergedValue:m,checked:h,mergedDisabled:n,cssVars:a?void 0:B,themeClass:null==E?void 0:E.themeClass,onRender:null==E?void 0:E.onRender}},render(){let{mergedClsPrefix:e,mergedDisabled:t,checked:a,mergedRailStyle:i,onRender:l,$slots:o}=this;null==l||l();let{checked:r,unchecked:n,icon:s,"checked-icon":u,"unchecked-icon":c}=o,p=!((0,y.aD)(s)&&(0,y.aD)(u)&&(0,y.aD)(c));return(0,d.h)("div",{role:"switch","aria-checked":a,class:[`${e}-switch`,this.themeClass,p&&`${e}-switch--icon`,a&&`${e}-switch--active`,t&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},(0,d.h)("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:i},(0,y.K9)(r,t=>(0,y.K9)(n,a=>t||a?(0,d.h)("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},(0,d.h)("div",{class:`${e}-switch__rail-placeholder`},(0,d.h)("div",{class:`${e}-switch__button-placeholder`}),t),(0,d.h)("div",{class:`${e}-switch__rail-placeholder`},(0,d.h)("div",{class:`${e}-switch__button-placeholder`}),a)):null)),(0,d.h)("div",{class:`${e}-switch__button`},(0,y.K9)(s,t=>(0,y.K9)(u,a=>(0,y.K9)(c,i=>(0,d.h)(m.Z,null,{default:()=>this.loading?(0,d.h)(h.Z,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(a||t)?(0,d.h)("div",{class:`${e}-switch__button-icon`,key:a?"checked-icon":"icon"},a||t):!this.checked&&(i||t)?(0,d.h)("div",{class:`${e}-switch__button-icon`,key:i?"unchecked-icon":"icon"},i||t):null})))),(0,y.K9)(r,t=>t&&(0,d.h)("div",{key:"checked",class:`${e}-switch__checked`},t)),(0,y.K9)(n,t=>t&&(0,d.h)("div",{key:"unchecked",class:`${e}-switch__unchecked`},t)))))}});var $=a(9506),C=a(825),z=a(3312),M=a(3776),F=a(6949),T=a(4203);let{t:R}=a(4958).Z.global,V=e=>T.e.get("/mailbox/list",{params:e}),A=e=>T.e.post("/mailbox/create",e,{fetchOptions:{loading:R("mailbox.api.loading.creating"),successMessage:!0}}),P=e=>T.e.post("/mailbox/update",e,{fetchOptions:{loading:R("mailbox.api.loading.updating"),successMessage:!0}}),q=e=>T.e.post("/mailbox/delete",e,{fetchOptions:{loading:R("mailbox.api.loading.deleting"),successMessage:!0}});var H=a(5476);let j={class:"flex"},K={class:"w-0 flex-1 flex items-center"},L={class:"min-w-0"},O=(0,d.aZ)({__name:"index",props:{value:{default:""}},setup(e){let{copyText:t}=(0,M.F)(),a=(0,d.iH)(!1),i=(0,d.Fl)(()=>a.value?e.value:"**********"),l=()=>{a.value=!a.value},o=async()=>{t(`${e.value}`)};return(e,t)=>{let r=H.ZP;return(0,d.wg)(),(0,d.iD)("div",j,[(0,d._)("div",K,[(0,d._)("div",L,[(0,d.Wm)(r,{class:"max-w-full",tooltip:{contentStyle:{maxWidth:"200px"}}},{default:(0,d.w5)(()=>[(0,d.Uk)((0,d.zw)((0,d.SU)(i)),1)]),_:1})]),(0,d._)("div",{class:"ml-6px text-15px cursor-pointer",title:"Show",onClick:l},[(0,d._)("i",{class:(0,d.C_)((0,d.SU)(a)?"i-mdi-eye-outline":"i-mdi-eye-off-outline")},null,2)]),(0,d._)("div",{class:"reset ml-6px text-15px cursor-pointer text-#333",title:"Copy",onClick:o},t[0]||(t[0]=[(0,d._)("i",{class:"i-mdi-content-copy"},null,-1)]))])])}}});var Q=a(5891),D=a(55),G=a(5778);let N=(0,d.aZ)({__name:"DomainSelect",props:(0,d.Vf)({isAll:{type:Boolean,default:!0}},{value:{default:()=>null},valueModifiers:{}}),emits:["update:value"],setup(e){let{t}=(0,G.QT)(),a=(0,d.tT)(e,"value"),i=(0,d.iH)(!1),l=(0,d.iH)([]);return(async()=>{try{i.value=!0;let o=await (0,D.LC)();(0,C.kJ)(o)&&(l.value=o.map(e=>({label:e.domain,value:e.domain})),e.isAll?(l.value.unshift({label:t("common.all.text"),value:""}),a.value=""):o.length>0&&!a.value&&(a.value=o[0].domain))}finally{i.value=!1}})(),(e,o)=>{let r=Q.Z;return(0,d.wg)(),(0,d.j4)(r,{value:a.value,"onUpdate:value":o[0]||(o[0]=e=>a.value=e),loading:(0,d.SU)(i),filterable:!0,options:(0,d.SU)(l),placeholder:(0,d.SU)(t)("mailbox.domain.selectPlaceholder")},null,8,["value","loading","options","placeholder"])}}}),I=(0,a(4222).default)(N,[["__scopeId","data-v-59ecdb27"]]);var X=a(5523),Y=a(4603),J=a(6883),ee=a(9167);let et=(0,x.cB)("input-group",`
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
 `)])])])])])]),ea=(0,d.aZ)({name:"InputGroup",props:{},setup(e){let{mergedClsPrefixRef:t}=(0,v.ZP)(e);return(0,ee.Z)("-input-group",et,t),{mergedClsPrefix:t}},render(){let{mergedClsPrefix:e}=this;return(0,d.h)("div",{class:`${e}-input-group`},this.$slots)}}),ei={class:"w-280px"},el={class:"w-280px"},eo={class:"w-280px"},er={class:"w-280px"},en=(0,d.aZ)({__name:"MailboxForm",setup(e){let{t}=(0,G.QT)(),a=(0,d.iH)(!1),i=(0,d.Fl)(()=>a.value?t("mailbox.form.editTitle"):t("mailbox.form.addTitle")),l=(0,d.AE)("formRef"),o=(0,d.qj)({quota:5,unit:"GB",isAdmin:0,full_name:"",domain:null,password:"",active:1}),r=[{label:"GB",value:"GB"},{label:"MB",value:"MB"}],s=[{label:t("mailbox.userType.general"),value:0},{label:t("mailbox.userType.admin"),value:1}],u={full_name:{trigger:"blur",validator:()=>""!==o.full_name.trim()&&!!o.domain||Error(t("mailbox.validation.emailRequired"))},password:{trigger:"blur",validator:()=>{if(a.value){if(o.password&&o.password.trim().length<8)return Error(t("mailbox.validation.passwordLength"));if(o.password&&!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/.test(o.password))return Error(t("mailbox.validation.passwordFormat"))}else{if(o.password.trim().length<8)return Error(t("mailbox.validation.passwordLength"));if(!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/.test(o.password))return Error(t("mailbox.validation.passwordFormat"))}return!0}}},c=(e,t)=>{switch(t){case"GB":e*=1024;case"MB":e*=1048576}return e},p=()=>({full_name:o.full_name,domain:o.domain||"",password:o.password,quota:c(o.quota,o.unit),isAdmin:o.isAdmin,active:o.active}),[m,h]=(0,z.d)({onChangeState:e=>{if(e){let e=h.getState(),{row:t}=e;if(a.value=e.isEdit,t){o.full_name=t.full_name,o.domain=t.domain,o.isAdmin=t.is_admin,o.active=t.active;let[e,a]=(0,C.UO)(t.quota).split(" ");o.quota=(0,C.Dx)(e),o.unit=a}}else o.full_name="",o.domain=null,o.password="",o.quota=5,o.unit="GB",o.isAdmin=0,o.active=1},onConfirm:async()=>{var e;await (null===(e=l.value)||void 0===e?void 0:e.validate());let t=p();a.value?await P(t):await A(t),h.getState().refresh()}});return(e,c)=>{let p=n.Z,h=J.ZP,b=Y.Z,v=Q.Z,f=X.Z;return(0,d.wg)(),(0,d.j4)((0,d.SU)(m),{title:(0,d.SU)(i),width:"520"},{default:(0,d.w5)(()=>[(0,d.Wm)(f,{ref_key:"formRef",ref:l,class:"pt-20px",model:(0,d.SU)(o),rules:u},{default:(0,d.w5)(()=>[(0,d.Wm)(h,{label:(0,d.SU)(t)("mailbox.form.emailAddress"),path:"full_name"},{default:(0,d.w5)(()=>[(0,d._)("div",ei,[(0,d.Wm)(ea,null,{default:(0,d.w5)(()=>[(0,d.Wm)(p,{value:(0,d.SU)(o).full_name,"onUpdate:value":c[0]||(c[0]=e=>(0,d.SU)(o).full_name=e),class:"flex-1",disabled:(0,d.SU)(a)},null,8,["value","disabled"]),(0,d.Wm)(I,{value:(0,d.SU)(o).domain,"onUpdate:value":c[1]||(c[1]=e=>(0,d.SU)(o).domain=e),class:"flex-1","is-all":!1,disabled:(0,d.SU)(a)},null,8,["value","disabled"])]),_:1})])]),_:1},8,["label"]),(0,d.Wm)(h,{label:(0,d.SU)(t)("mailbox.form.password"),path:"password"},{default:(0,d.w5)(()=>[(0,d._)("div",el,[(0,d.Wm)(p,{value:(0,d.SU)(o).password,"onUpdate:value":c[2]||(c[2]=e=>(0,d.SU)(o).password=e),placeholder:(0,d.SU)(t)("mailbox.form.passwordPlaceholder")},null,8,["value","placeholder"])])]),_:1},8,["label"]),(0,d.Wm)(h,{label:(0,d.SU)(t)("mailbox.form.quota"),path:"quota"},{default:(0,d.w5)(()=>[(0,d._)("div",eo,[(0,d.Wm)(ea,null,{default:(0,d.w5)(()=>[(0,d.Wm)(b,{value:(0,d.SU)(o).quota,"onUpdate:value":c[3]||(c[3]=e=>(0,d.SU)(o).quota=e),class:"flex-1",min:1,"show-button":!1},null,8,["value"]),(0,d.Wm)(v,{value:(0,d.SU)(o).unit,"onUpdate:value":c[4]||(c[4]=e=>(0,d.SU)(o).unit=e),class:"flex-1",options:r},null,8,["value"])]),_:1})])]),_:1},8,["label"]),(0,d.Wm)(h,{label:(0,d.SU)(t)("mailbox.form.userType"),path:"userType"},{default:(0,d.w5)(()=>[(0,d._)("div",er,[(0,d.Wm)(v,{value:(0,d.SU)(o).isAdmin,"onUpdate:value":c[5]||(c[5]=e=>(0,d.SU)(o).isAdmin=e),options:s},null,8,["value"])])]),_:1},8,["label"]),(0,d.Wm)(h,{label:(0,d.SU)(t)("mailbox.form.status")},{default:(0,d.w5)(()=>[(0,d.Wm)(Z,{value:(0,d.SU)(o).active,"onUpdate:value":c[6]||(c[6]=e=>(0,d.SU)(o).active=e),"checked-value":1,"unchecked-value":0},null,8,["value"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["title"])}}}),ed={class:"p-24px"},es={class:"mb-20px text-24px font-bold"},eu={class:"w-220px"};function ec(e){return"function"==typeof e||"[object Object]"===Object.prototype.toString.call(e)&&!(0,d.lA)(e)}let ep=(0,d.aZ)({__name:"index",setup(e){let{t}=(0,G.QT)(),{copyText:a}=(0,M.F)(),{tableParams:i,tableList:n,loading:c,tableTotal:p,getTableData:m}=(0,F.Z)({loading:!0,immediate:!1,params:{page:1,page_size:10,domain:null,keyword:""},fetchFn:V}),h=(0,d.iH)([{key:"username",title:t("mailbox.columns.username"),minWidth:120,ellipsis:{tooltip:!0}},{key:"password",title:t("mailbox.columns.password"),width:"18%",minWidth:120,render:e=>(0,d.Wm)(O,{value:e.password||"--"},null)},{key:"login",title:t("mailbox.columns.loginInfo"),render:e=>{let i;return(0,d.Wm)(u.ZP,{text:!0,type:"primary",onClick:()=>{a(t("mailbox.loginInfo.template",{webmail:window.location.origin+"/roundcube",username:e.username,password:e.password}))}},ec(i=t("common.actions.copy"))?i:{default:()=>[i]})}},{key:"quota",title:t("mailbox.columns.quota"),width:"14%",minWidth:120,render:e=>(0,C.UO)(e.quota)},{key:"is_admin",title:t("mailbox.columns.type"),width:"14%",minWidth:100,render:e=>1===e.is_admin?t("mailbox.userType.admin"):t("mailbox.userType.general")},{key:"status",title:t("mailbox.columns.status"),width:"12%",minWidth:80,render:e=>(0,d.Wm)(Z,{value:e.active,"checked-value":1,"unchecked-value":0,size:"small",onUpdateValue:t=>{g(e,t)}},null)},{title:t("common.columns.actions"),key:"actions",align:"right",width:120,render:e=>{let a,i;return(0,d.Wm)($.Z,{inline:!0},{default:()=>[(0,d.Wm)(u.ZP,{type:"primary",text:!0,onClick:()=>{w(e)}},ec(a=t("common.actions.edit"))?a:{default:()=>[a]}),(0,d.Wm)(u.ZP,{type:"error",text:!0,onClick:()=>{x(e)}},ec(i=t("common.actions.delete"))?i:{default:()=>[i]})]})}}]),[b,v]=(0,z.d)({component:en,state:{isEdit:!1,refresh:m}}),f=()=>{v.setState({isEdit:!1,row:null}),v.open()},g=async(e,t)=>{await P({full_name:e.full_name,domain:e.domain,password:e.password,quota:e.quota,isAdmin:e.is_admin,active:t}),e.active=t},w=e=>{v.setState({isEdit:!0,row:e}),v.open()},x=e=>{(0,C.iG)({title:t("mailbox.delete.title"),content:t("mailbox.delete.confirm",{name:e.username}),confirmText:t("common.actions.delete"),confirmType:"error",onConfirm:async()=>{await q({email:e.username}),m()}})};return(e,a)=>{let v=r.Z,g=o.Z,w=l.Z;return(0,d.wg)(),(0,d.iD)("div",ed,[(0,d._)("div",es,(0,d.zw)((0,d.SU)(t)("mailbox.title")),1),(0,d.Wm)(w,null,{toolsLeft:(0,d.w5)(()=>[(0,d.Wm)((0,d.SU)(u.ZP),{type:"primary",onClick:f},{default:(0,d.w5)(()=>[(0,d.Uk)((0,d.zw)((0,d.SU)(t)("mailbox.actions.add")),1)]),_:1})]),toolsRight:(0,d.w5)(()=>[(0,d._)("div",eu,[(0,d.Wm)(I,{value:(0,d.SU)(i).domain,"onUpdate:value":[a[0]||(a[0]=e=>(0,d.SU)(i).domain=e),a[1]||(a[1]=()=>(0,d.SU)(m)(!0))]},null,8,["value"])]),(0,d.Wm)(s,{value:(0,d.SU)(i).keyword,"onUpdate:value":a[2]||(a[2]=e=>(0,d.SU)(i).keyword=e),width:"280",placeholder:(0,d.SU)(t)("mailbox.search.usernamePlaceholder"),onSearch:a[3]||(a[3]=()=>(0,d.SU)(m)(!0))},null,8,["value","placeholder"])]),table:(0,d.w5)(()=>[(0,d.Wm)(v,{loading:(0,d.SU)(c),columns:(0,d.SU)(h),data:(0,d.SU)(n)},null,8,["loading","columns","data"])]),pageRight:(0,d.w5)(()=>[(0,d.Wm)(g,{page:(0,d.SU)(i).page,"onUpdate:page":a[4]||(a[4]=e=>(0,d.SU)(i).page=e),"page-size":(0,d.SU)(i).page_size,"onUpdate:pageSize":a[5]||(a[5]=e=>(0,d.SU)(i).page_size=e),"item-count":(0,d.SU)(p),onRefresh:a[6]||(a[6]=()=>(0,d.SU)(m)(!0))},null,8,["page","page-size","item-count"])]),modal:(0,d.w5)(()=>[(0,d.Wm)((0,d.SU)(b))]),_:1})])}}})}}]);