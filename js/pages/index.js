import { header } from "../layout/header.js";
import { leftBloc,middleBloc,rigthtBloc } from "../layout/corps.js";
import {  body, Header,corpsContent,menuGauche,menuMilieu,menuDroite } from "../helper/bigBlocContent.js";

const indexPage=(data)=>{
    console.log(data);
    // console.log("data : "+ data["Datas"][0]["Posts"]["Content"]);

   let posts=data["Datas"] ?  data["Datas"] : ""; 

    let CatgoryArray=data["Category"];
    let Username=data["User"]["Username"];
    body.appendChild(Header);
    corpsContent.appendChild(menuGauche);
    corpsContent.appendChild(menuMilieu);
    corpsContent.appendChild(menuDroite);
    body.appendChild(corpsContent);

    const _header=document.querySelector('header');
    // const _corpsContent=document.querySelector('.corps');
    const _leftBloc=document.querySelector('.menu-gauche');
    const _middleBloc=document.querySelector('.milieu');
    const _rigthtBloc=document.querySelector('.menu-droite');

    header(_header,Username);
    leftBloc(_leftBloc,CatgoryArray);
    middleBloc(_middleBloc,posts);
    rigthtBloc(_rigthtBloc);
    
}

export {indexPage}