import { header } from "../layout/header.js";
import { leftBloc,middleBloc,rigthtBloc } from "../layout/corps.js";
import {  body, Header,corpsContent,menuGauche,menuMilieu,menuDroite } from "../helper/bigBlocContent.js";

const indexPage=(data)=>{
    console.log(data);
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

    header(_header)
    leftBloc(_leftBloc);
    middleBloc(_middleBloc);
    rigthtBloc(_rigthtBloc);
}

export {indexPage}