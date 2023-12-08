import {  routes,addRouter,replaceRouter,currentPath ,titlePage} from "./router/route.js";


import { linkApi } from "./helper/api_link.js";        
import { isSessionFoundBoolean } from "./verify_sessions.js";
import { displayFom } from "./pages/signUpSignIn.js";
import {signInForm ,signUpForm} from "./auth/forms.js";
import { indexPage } from "./pages/index.js";
import { page404 } from "./pages/page404.js";

import { filterPost } from "./post/filterPost.js";
import { postsFilter } from "./post/postsFiltered.js";
import { displayFormMecanisme,readMore ,displayFomPost,seeMore,disPlayComment,disPlayCommentFilter} from "./helper/diplayingMecanisme.js";
import { addPostFrom } from "./pages/addPostForm.js";
import { formAddPost } from "./post/formAddPost.js";


import { displayComment } from "./layout/corps.js";
import { getComments } from "./comment/getComment.js";
import { addComment } from "./comment/addComment.js";

import { liskePost } from "./likeDislike/post/like.js";
import { DisLiskePost } from "./likeDislike/post/dislike.js";
import { liskeComment } from "./likeDislike/comment/like.js";
import { DisLiskeComment } from "./likeDislike/comment/dislike.js";





import {logOut} from "./auth/logOut.js";



const main=()=>{ 
    
    document.addEventListener('DOMContentLoaded', (event) => {
        // 404 page if route is note correcte
     let rd=routes[currentPath];       
      if (!rd) {
        document.querySelector('body').innerHTML=page404();
        console.error("404 Page not found");
        return
      }

    async function checkUserSession() {
        const data = await isSessionFoundBoolean();
        if (data===undefined) {           
            Authentification() ;     
        }else{

             // make route
            titlePage("Home");
            let r=routes["/Home"]['name'];
            replaceRouter(r);

            //afficher l'interface des posts
            indexPage(data); 


            // add post fom et  methode
            const myModal=document.querySelector('#myModal');
            const btn = document.querySelector(".create-post-btn");
            // creer le formulaire de creation de post et le cacher
            addPostFrom(myModal,data);
            
            const span = document.getElementsByClassName("close")[0];
            // affciher le formulaire de creation de post
            displayFomPost(btn,myModal,span)

           const formCreatPost= document.querySelector('.form-creat-post');
            //    ajouter une post
           formAddPost(formCreatPost);


            const readMoreButton=document.querySelectorAll('.myBtn');
            const cardDescription = document.querySelectorAll(".card-description");
            const onePostBlocks=document.querySelectorAll(".one-post-block");
            // tronc long text post
            seeMore(cardDescription,readMoreButton);                  
            readMore(cardDescription,readMoreButton,onePostBlocks);
                
            // log out 
            let btnLogOut=document.querySelector('.log-out');
            logOut(btnLogOut);

            let comments = document.querySelectorAll('.comment');
            let IdPost = document.querySelectorAll('.Id-post');
            let createCommentForm=document.querySelectorAll('.create-comment');
            let blocComment=document.querySelectorAll('.bloc-comment');                
            let lastPost=document.querySelector('.last-post');
            let lastFormComment=document.querySelector('.last-form-comment');
            let lastBlocComment=document.querySelector('.last-bloc-comment');

            // gerer l'affichage des commenataires
            disPlayComment(comments,createCommentForm,lastPost,lastFormComment,lastBlocComment);

            // recuperer et affcher les  commentaires
            getComments(comments, IdPost,createCommentForm, function(_data) {                        
                displayComment(blocComment,_data,createCommentForm,lastBlocComment);                        
               
                // like an dislike Comments
                const likeComment=document.querySelectorAll('.like-comment');
                const dislikeComment=document.querySelectorAll('.dislike-comment');
                const likeCommentId=document.querySelectorAll('.id-comment-like');
                const dislikeCommentId=document.querySelectorAll('.id-comment-dislike');
                const likeCommentScore=document.querySelectorAll('.scorecommentLike');
                const dislikeCommentScore=document.querySelectorAll('.scorecommentDisLike');
    
                liskeComment(likeComment,likeCommentId,likeCommentScore,dislikeCommentScore,dislikeComment);
                DisLiskeComment(dislikeComment,dislikeCommentId,likeCommentScore,dislikeCommentScore,likeComment);
            });

            // ajouter un commentair
            const formComment=document.querySelectorAll('.form-comment');
            addComment(formComment);

            // like an dislike Posts
            const likePost=document.querySelectorAll('.like-post');
            const dislikePost=document.querySelectorAll('.dislike-post');
            const likePostId=document.querySelectorAll('.id-post-like');
            const dislikePostId=document.querySelectorAll('.id-post-dislike');
            const likePostScore=document.querySelectorAll('.scoreLike');
            const dislikePostScore=document.querySelectorAll('.scoreDisLike');
  
            liskePost(likePost,likePostId,likePostScore,dislikePostScore,dislikePost);
            DisLiskePost(dislikePost,dislikePostId,likePostScore,dislikePostScore,likePost);

            // filter category diplaying
            filterByCategory();
         

        }
    }

    (async () => {
        await checkUserSession();
    })();
  


}); 

}

function Authentification() {
      // make route
      titlePage("Authentification");
      let r=routes["/Login"]['name'];
      replaceRouter(r);
      
      // afficher les formualire d'authentifications
       displayFom();

      const formSignIn=document.querySelector('.form-1');
      const formSignUp=document.querySelector('.form-2');
      const _ContentForms=document.querySelector('.content');
      let spinner = document.querySelector('.spinner');
      let row=document.querySelectorAll('.row');

      // gerer la navigation entre les deux formulaire
      displayFormMecanisme(formSignUp,formSignIn,row);
      
      // enoyer les donnes du formulaire
      if(formSignIn) signInForm(_ContentForms,formSignIn,formSignUp,spinner,linkApi)
      if(formSignUp) signUpForm(_ContentForms,formSignUp,formSignIn,spinner,linkApi)
}


// filter posts
function filterByCategory() {
                const contenCatId=document.querySelectorAll('.contenCatId');
                const categoryId=document.querySelectorAll('.categoryId');
                const contentPostBlock=document.querySelector('.content-post-block');
    
               filterPost(contenCatId,categoryId, function(_data) {                 
                    // nettoyer et avec les posts filtre
                    contentPostBlock.innerHTML=""
                    let div= document.createElement('div');
                    div.innerHTML= _data ? postsFilter(_data): contentPostBlock.innerHTML="<p id='err'>NO POST FOUND </p>"
                  
                    contentPostBlock.appendChild(div);
    
                    let _comments=document.querySelectorAll('.one-post-block .post-content .content-poster-like .content-poster .like-comment-block .comment')
                    let _createCommentForm=document.querySelectorAll('.create-comment');
                    let _blocComment=document.querySelectorAll('.bloc-comment');                
                    let _lastPost=document.querySelector('.last-post');
                    let _lastFormComment=document.querySelector('.last-form-comment');
                    let _lastBlocComment=document.querySelector('.last-bloc-comment');
                 
                    disPlayCommentFilter(_comments,_createCommentForm,_blocComment,_lastPost,_lastFormComment,_lastBlocComment);
                    
                    const _readMoreButton=document.querySelectorAll('.one-post-block .post-content .myBtn');
                    const _cardDescription = document.querySelectorAll(".one-post-block .post-content .post-text .card-description");
                    const _onePostBlocks=document.querySelectorAll(".one-post-block");
                    // tronc long text post
                    seeMore(_cardDescription,_readMoreButton);                  
                    readMore(_cardDescription,_readMoreButton,_onePostBlocks);
                    const _formComment=document.querySelectorAll('.create-comment .form-comment');
                    addComment(_formComment);
                  
                    if (contentPostBlock.textContent==="") contentPostBlock.innerHTML="<p id='err'>NO POST FOUND </p>"
    
                    // like an dislike
                    const likePost=document.querySelectorAll('.like-post');
                    const dislikePost=document.querySelectorAll('.dislike-post');
                    const likePostId=document.querySelectorAll('.id-post-like');
                    const dislikePostId=document.querySelectorAll('.id-post-dislike');
                    const likePostScore=document.querySelectorAll('.scoreLike');
                    const dislikePostScore=document.querySelectorAll('.scoreDisLike');
        
                    liskePost(likePost,likePostId,likePostScore,dislikePostScore,dislikePost);
                    DisLiskePost(dislikePost,dislikePostId,likePostScore,dislikePostScore,likePost);
    


                    // like an dislike Comments
                    const likeComment=document.querySelectorAll('.like-comment');
                    const dislikeComment=document.querySelectorAll('.dislike-comment');
                    const likeCommentId=document.querySelectorAll('.id-comment-like');
                    const dislikeCommentId=document.querySelectorAll('.id-comment-dislike');
                    const likeCommentScore=document.querySelectorAll('.scorecommentLike');
                    const dislikeCommentScore=document.querySelectorAll('.scorecommentDisLike');
    
                    liskeComment(likeComment,likeCommentId,likeCommentScore,dislikeCommentScore,dislikeComment);
                    DisLiskeComment(dislikeComment,dislikeCommentId,likeCommentScore,dislikeCommentScore,likeComment);
    
    
                 });
}

export{
    main
}


