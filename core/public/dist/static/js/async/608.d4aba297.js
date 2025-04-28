"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["608"],{4944:function(e,t,n){n.d(t,{L:()=>i});function i(e,t){if(!e)return;let n=document.createElement("a");n.href=e,void 0!==t&&(n.download=t),document.body.appendChild(n),n.click(),document.body.removeChild(n)}},3012:function(e,t,n){n.d(t,{W3:()=>p,ZP:()=>f});var i=n(209),a=n(1321),l=n(4124),r=n(6169),o=n(1579),s=n(8755);let c={fontWeightActive:"400"},d={name:"Breadcrumb",common:s.Z,self:function(e){let{fontSize:t,textColor3:n,textColor2:i,borderRadius:a,buttonColor2Hover:l,buttonColor2Pressed:r}=e;return Object.assign(Object.assign({},c),{fontSize:t,itemLineHeight:"1.25",itemTextColor:n,itemTextColorHover:i,itemTextColorPressed:i,itemTextColorActive:i,itemBorderRadius:a,itemColorHover:l,itemColorPressed:r,separatorColor:n})}};var u=n(2249);let m=(0,u.cB)("breadcrumb",`
 white-space: nowrap;
 cursor: default;
 line-height: var(--n-item-line-height);
`,[(0,u.c)("ul",`
 list-style: none;
 padding: 0;
 margin: 0;
 `),(0,u.c)("a",`
 color: inherit;
 text-decoration: inherit;
 `),(0,u.cB)("breadcrumb-item",`
 font-size: var(--n-font-size);
 transition: color .3s var(--n-bezier);
 display: inline-flex;
 align-items: center;
 `,[(0,u.cB)("icon",`
 font-size: 18px;
 vertical-align: -.2em;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `),(0,u.c)("&:not(:last-child)",[(0,u.cM)("clickable",[(0,u.cE)("link",`
 cursor: pointer;
 `,[(0,u.c)("&:hover",`
 background-color: var(--n-item-color-hover);
 `),(0,u.c)("&:active",`
 background-color: var(--n-item-color-pressed); 
 `)])])]),(0,u.cE)("link",`
 padding: 4px;
 border-radius: var(--n-item-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 position: relative;
 `,[(0,u.c)("&:hover",`
 color: var(--n-item-text-color-hover);
 `,[(0,u.cB)("icon",`
 color: var(--n-item-text-color-hover);
 `)]),(0,u.c)("&:active",`
 color: var(--n-item-text-color-pressed);
 `,[(0,u.cB)("icon",`
 color: var(--n-item-text-color-pressed);
 `)])]),(0,u.cE)("separator",`
 margin: 0 8px;
 color: var(--n-separator-color);
 transition: color .3s var(--n-bezier);
 user-select: none;
 -webkit-user-select: none;
 `),(0,u.c)("&:last-child",[(0,u.cE)("link",`
 font-weight: var(--n-font-weight-active);
 cursor: unset;
 color: var(--n-item-text-color-active);
 `,[(0,u.cB)("icon",`
 color: var(--n-item-text-color-active);
 `)]),(0,u.cE)("separator",`
 display: none;
 `)])])]),p=(0,o.U)("n-breadcrumb"),v=Object.assign(Object.assign({},a.Z.props),{separator:{type:String,default:"/"}}),f=(0,i.aZ)({name:"Breadcrumb",props:v,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:n}=(0,l.ZP)(e),o=(0,a.Z)("Breadcrumb","-breadcrumb",m,d,e,t);(0,i.JJ)(p,{separatorRef:(0,i.Vh)(e,"separator"),mergedClsPrefixRef:t});let s=(0,i.Fl)(()=>{let{common:{cubicBezierEaseInOut:e},self:{separatorColor:t,itemTextColor:n,itemTextColorHover:i,itemTextColorPressed:a,itemTextColorActive:l,fontSize:r,fontWeightActive:s,itemBorderRadius:c,itemColorHover:d,itemColorPressed:u,itemLineHeight:m}}=o.value;return{"--n-font-size":r,"--n-bezier":e,"--n-item-text-color":n,"--n-item-text-color-hover":i,"--n-item-text-color-pressed":a,"--n-item-text-color-active":l,"--n-separator-color":t,"--n-item-color-hover":d,"--n-item-color-pressed":u,"--n-item-border-radius":c,"--n-font-weight-active":s,"--n-item-line-height":m}}),c=n?(0,r.F)("breadcrumb",void 0,s,e):void 0;return{mergedClsPrefix:t,cssVars:n?void 0:s,themeClass:null==c?void 0:c.themeClass,onRender:null==c?void 0:c.onRender}},render(){var e;return null===(e=this.onRender)||void 0===e||e.call(this),(0,i.h)("nav",{class:[`${this.mergedClsPrefix}-breadcrumb`,this.themeClass],style:this.cssVars,"aria-label":"Breadcrumb"},(0,i.h)("ul",null,this.$slots))}})},6666:function(e,t,n){n.d(t,{Z:()=>o});var i=n(209),a=n(8282),l=n(4934),r=n(3012);let o=(0,i.aZ)({name:"BreadcrumbItem",props:{separator:String,href:String,clickable:{type:Boolean,default:!0},onClick:Function},slots:Object,setup(e,{slots:t}){let n=(0,i.f3)(r.W3,null);if(!n)return()=>null;let{separatorRef:o,mergedClsPrefixRef:s}=n,c=function(e=l.j?window:null){let t=()=>{let{hash:t,host:n,hostname:i,href:a,origin:l,pathname:r,port:o,protocol:s,search:c}=(null==e?void 0:e.location)||{};return{hash:t,host:n,hostname:i,href:a,origin:l,pathname:r,port:o,protocol:s,search:c}},n=(0,i.iH)(t()),a=()=>{n.value=t()};return(0,i.bv)(()=>{e&&(e.addEventListener("popstate",a),e.addEventListener("hashchange",a))}),(0,i.SK)(()=>{e&&(e.removeEventListener("popstate",a),e.removeEventListener("hashchange",a))}),n}(),d=(0,i.Fl)(()=>e.href?"a":"span"),u=(0,i.Fl)(()=>c.value.href===e.href?"location":null);return()=>{let{value:n}=s;return(0,i.h)("li",{class:[`${n}-breadcrumb-item`,e.clickable&&`${n}-breadcrumb-item--clickable`]},(0,i.h)(d.value,{class:`${n}-breadcrumb-item__link`,"aria-current":u.value,href:e.href,onClick:e.onClick},t),(0,i.h)("span",{class:`${n}-breadcrumb-item__separator`,"aria-hidden":"true"},(0,a.gI)(t.separator,()=>{var t;return[null!==(t=e.separator)&&void 0!==t?t:o.value]})))}}})},8951:function(e,t,n){n.d(t,{Z:()=>h});var i=n(5083),a=n(2370),l=n(209),r=n(4131),o=n(1321),s=n(4124),c=n(6169),d=n(2249);let u={name:"Spin",common:n(8755).Z,self:function(e){let{opacityDisabled:t,heightTiny:n,heightSmall:i,heightMedium:a,heightLarge:l,heightHuge:r,primaryColor:o,fontSize:s}=e;return{fontSize:s,textColor:o,sizeTiny:n,sizeSmall:i,sizeMedium:a,sizeLarge:l,sizeHuge:r,color:o,opacitySpinning:t}}};var m=n(3291);let p=(0,d.c)([(0,d.c)("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),(0,d.cB)("spin-container",`
 position: relative;
 `,[(0,d.cB)("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[(0,m.h)()])]),(0,d.cB)("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),(0,d.cB)("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[(0,d.cM)("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),(0,d.cB)("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),(0,d.cB)("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[(0,d.cM)("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),v={small:20,medium:18,large:16},f=Object.assign(Object.assign({},o.Z.props),{contentClass:String,contentStyle:[Object,String],description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number}),h=(0,l.aZ)({name:"Spin",props:f,slots:Object,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:n}=(0,s.ZP)(e),r=(0,o.Z)("Spin","-spin",p,u,e,t),m=(0,l.Fl)(()=>{let{size:t}=e,{common:{cubicBezierEaseInOut:n},self:a}=r.value,{opacitySpinning:l,color:o,textColor:s}=a;return{"--n-bezier":n,"--n-opacity-spinning":l,"--n-size":"number"==typeof t?(0,i.BL)(t):a[(0,d.Tl)("size",t)],"--n-color":o,"--n-text-color":s}}),f=n?(0,c.F)("spin",(0,l.Fl)(()=>{let{size:t}=e;return"number"==typeof t?String(t):t[0]}),m,e):void 0,h=(0,a.Z)(e,["spinning","show"]),g=(0,l.iH)(!1);return(0,l.m0)(t=>{let n;if(h.value){let{delay:i}=e;if(i){n=window.setTimeout(()=>{g.value=!0},i),t(()=>{clearTimeout(n)});return}}g.value=h.value}),{mergedClsPrefix:t,active:g,mergedStrokeWidth:(0,l.Fl)(()=>{let{strokeWidth:t}=e;if(void 0!==t)return t;let{size:n}=e;return v["number"==typeof n?"medium":n]}),cssVars:n?void 0:m,themeClass:null==f?void 0:f.themeClass,onRender:null==f?void 0:f.onRender}},render(){var e,t;let{$slots:n,mergedClsPrefix:i,description:a}=this,o=n.icon&&this.rotate,s=(a||n.description)&&(0,l.h)("div",{class:`${i}-spin-description`},a||(null===(e=n.description)||void 0===e?void 0:e.call(n))),c=n.icon?(0,l.h)("div",{class:[`${i}-spin-body`,this.themeClass]},(0,l.h)("div",{class:[`${i}-spin`,o&&`${i}-spin--rotate`],style:n.default?"":this.cssVars},n.icon()),s):(0,l.h)("div",{class:[`${i}-spin-body`,this.themeClass]},(0,l.h)(r.Z,{clsPrefix:i,style:n.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${i}-spin`}),s);return null===(t=this.onRender)||void 0===t||t.call(this),n.default?(0,l.h)("div",{class:[`${i}-spin-container`,this.themeClass],style:this.cssVars},(0,l.h)("div",{class:[`${i}-spin-content`,this.active&&`${i}-spin-content--spinning`,this.contentClass],style:this.contentStyle},n),(0,l.h)(l.uT,{name:"fade-in-transition"},{default:()=>this.active?c:null})):c}})},1555:function(e,t,n){n.d(t,{A8:()=>a,NI:()=>l,QB:()=>s,WF:()=>o,Xc:()=>d,ef:()=>c,ij:()=>r});var i=n(4203);function a(e){return i.e.get("/email_template/list",{params:e})}function l(){return i.e.get("/email_template/all")}function r(e){return i.e.get("/email_template/get",{params:e})}function o(e){return i.e.post("/email_template/create",e,{fetchOptions:{loading:"Creating template, please wait...",successMessage:!0}})}function s(e){return i.e.post("/email_template/update",e,{fetchOptions:{loading:"Updating template, please wait...",successMessage:!0}})}function c(e){return i.e.post("/email_template/copy",e,{fetchOptions:{loading:"Copying template, please wait...",successMessage:!0}})}function d(e){return i.e.post("/email_template/delete",e,{fetchOptions:{loading:"Deleting template, please wait...",successMessage:!0}})}},3104:function(e,t,n){n.d(t,{Z:()=>l});var i=n(209);let a=["height","srcdoc"],l=(0,i.aZ)({__name:"index",props:{value:{type:String,default:""},height:{type:[String,Number],default:"100%"}},setup:e=>(t,n)=>((0,i.wg)(),(0,i.iD)("iframe",{width:"100%",height:e.height,srcdoc:e.value,style:{border:"none"}},null,8,a))})},2329:function(e,t,n){n.r(t),n.d(t,{default:()=>$});var i=n(3447),a=n(5523),l=n(6528),r=n(3104),o=n(8951),s=n(209);let c=(0,s.aZ)({__name:"index",props:(0,s.Vf)({language:{type:String,default:"html"},theme:{type:String,default:"vs-dark"},options:{type:Object,default:()=>({})}},{value:{default:""},valueModifiers:{}}),emits:["update:value"],setup(e,t){let{expose:i}=t,a=(0,s.tT)(e,"value"),l=(0,s.iH)(null),r=null,c=null,d=(0,s.iH)(!1),u=async()=>{if(l.value&&!r&&!c)try{d.value=!0;let t=await n.e("491").then(n.bind(n,2238));(c=(r=await t.default.init()).editor.create(l.value,{value:a.value,language:e.language,theme:e.theme,minimap:{enabled:!1},automaticLayout:!0,scrollBeyondLastLine:!1,fontSize:14,tabSize:2,wordWrap:"on",...e.options})).onDidChangeModelContent(()=>{let e=(null==c?void 0:c.getValue())||"";a.value=e})}finally{d.value=!1}},m=()=>{c&&(c.dispose(),c=null)};return(0,s.YP)(()=>a.value,e=>{c&&e!==c.getValue()&&c.setValue(e)}),(0,s.bv)(()=>{u()}),(0,s.Jd)(()=>{m()}),i({getEditor:()=>c}),(e,t)=>{let n=o.Z;return(0,s.wg)(),(0,s.j4)(n,{class:"h-full",show:(0,s.SU)(d)},{default:(0,s.w5)(()=>[(0,s._)("div",{ref_key:"containerRef",ref:l,class:"monaco-editor-container"},null,512)]),_:1},8,["show"])}}});var d=n(4222);let u=(0,d.default)(c,[["__scopeId","data-v-099fb93b"]]);var m=n(919),p=n(6883),v=n(8454),f=n(3012),h=n(6666),g=n(825),b=n(1555),w=n(5e3);let y={class:"template-edit"},_={class:"w-400px"},k={class:"editor-area"},x={class:"mb-16px"},z={class:"flex-1 overflow-hidden"},C={class:"preview-title"},S={class:"flex-1 overflow-hidden"},Z={class:"action-buttons"},B=(0,s.aZ)({__name:"edit",setup(e){let t=(0,w.yj)(),n=(0,s.iH)(`${t.params.id}`),o=(0,s.qj)({id:0,temp_name:"",content:""}),c=e=>{let t=new FileReader;t.onload=e=>{var t;o.content=null===(t=e.target)||void 0===t?void 0:t.result},e.file&&t.readAsText(e.file)},d=async()=>{await (0,b.QB)({id:o.id,temp_name:o.temp_name,content:o.content})},B=()=>{window.history.back()};return(async()=>{let e=await (0,b.ij)({id:n.value});(0,g.Kn)(e)&&(o.id=e.id,o.temp_name=e.temp_name,o.content=e.content)})(),(e,t)=>{let n=(0,s.up)("router-link"),g=h.Z,b=f.ZP,w=v.Z,$=p.ZP,W=m.Z,j=r.Z,U=l.ZP,O=a.Z,P=i.ZP;return(0,s.wg)(),(0,s.iD)("div",y,[(0,s.Wm)(b,{class:"mb-20px"},{default:(0,s.w5)(()=>[(0,s.Wm)(g,null,{default:(0,s.w5)(()=>[(0,s.Wm)(n,{to:"/market/template"},{default:(0,s.w5)(()=>[(0,s.Uk)((0,s.zw)(e.$t("market.template.title")),1)]),_:1})]),_:1}),(0,s.Wm)(g,null,{default:(0,s.w5)(()=>[(0,s.Uk)((0,s.zw)((0,s.SU)(o).temp_name),1)]),_:1})]),_:1}),(0,s.Wm)(O,{class:"template-form","label-placement":"top"},{default:(0,s.w5)(()=>[(0,s.Wm)($,{label:e.$t("market.template.name")},{default:(0,s.w5)(()=>[(0,s._)("div",_,[(0,s.Wm)(w,{value:(0,s.SU)(o).temp_name,"onUpdate:value":t[0]||(t[0]=e=>(0,s.SU)(o).temp_name=e)},null,8,["value"])])]),_:1},8,["label"]),(0,s.Wm)($,{class:"editor-form-item",label:e.$t("market.template.content"),"show-feedback":!1},{default:(0,s.w5)(()=>[(0,s._)("div",k,[(0,s._)("div",x,[(0,s.Wm)(W,{mode:"button","is-upload":!1,accept:["html"],onChange:c},{default:(0,s.w5)(()=>[(0,s.Uk)((0,s.zw)(e.$t("market.template.uploadHtml")),1)]),_:1})]),(0,s._)("div",z,[(0,s.Wm)(u,{value:(0,s.SU)(o).content,"onUpdate:value":t[1]||(t[1]=e=>(0,s.SU)(o).content=e),language:"html"},null,8,["value"])])]),(0,s.Wm)(U,{class:"preview-area","content-style":{display:"flex",flexDirection:"column"}},{default:(0,s.w5)(()=>[(0,s._)("div",C,(0,s.zw)(e.$t("common.actions.preview")),1),(0,s._)("div",S,[(0,s.Wm)(j,{value:(0,s.SU)(o).content},null,8,["value"])])]),_:1})]),_:1},8,["label"])]),_:1}),(0,s._)("div",Z,[(0,s.Wm)(P,{type:"primary",onClick:d},{default:(0,s.w5)(()=>[(0,s.Uk)((0,s.zw)(e.$t("common.actions.save")),1)]),_:1}),(0,s.Wm)(P,{class:"ml-16px",onClick:B},{default:(0,s.w5)(()=>[(0,s.Uk)((0,s.zw)(e.$t("common.actions.back")),1)]),_:1})])])}}}),$=(0,d.default)(B,[["__scopeId","data-v-243d53e2"]])}}]);