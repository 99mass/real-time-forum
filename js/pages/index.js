import {  } from "../layout/corps.js";

import { header } from "../layout/header.js";
import { leftBloc,middleBloc,rigthtBloc } from "../layout/corps.js";




const _corps=document.querySelector('.corps');

const indexPage=()=>{
    
    header(true)
    leftBloc();
    middleBloc();
    rigthtBloc();
}

export {indexPage}