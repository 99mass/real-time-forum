
import {  body, Header,corpsContent } from "../helper/bigBlocContent.js";

body.appendChild(Header);
body.appendChild(corpsContent);

const _header=document.querySelector('header')

const header = (state)=>{

     _header.innerHTML=`
        <div class="titre">
        <img src="assets/forum-message-svgrepo-com.svg" alt="">
            <span>real time forum</span>
        </div>
        <div class="content-input"><input type="search" placeholder="Type here to search..."></div>
        <div class="btns">
            <div class="profile">
                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                <span>ssambadi</span>
            </div>
            <button>
                <span>LogOut</span>
                <img src="assets/logout-svgrepo-com.svg" alt="">
            </button>
        </div>
      `;

}

export {header}