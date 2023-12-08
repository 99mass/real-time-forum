import { linkApi } from "./helper/api_link.js";        
import { isSessionFoundBoolean } from "./verify_sessions.js";
import { displayFom } from "./pages/signUpSignIn.js";
import {signInForm ,signUpForm} from "./auth/forms.js";
import { indexPage } from "./pages/index.js";
import { filterPost } from "./post/filterPost.js";
import { postsFilter } from "./post/postsFiltered.js";
import { displayFormMecanisme,readMore ,displayFomPost,seeMore,disPlayComment,disPlayCommentFilter} from "./helper/diplayingMecanisme.js";
import { addPostFrom } from "./pages/addPostForm.js";
import { formAddPost } from "./post/formAddPost.js";


import { displayComment } from "./layout/corps.js";
import { getComments } from "./comment/getComment.js";
import { addComment } from "./comment/addComment.js";
import {logOut} from "./auth/logOut.js";



const main=()=>{ 
    
    document.addEventListener('DOMContentLoaded', (event) => {

    async function checkUserSession() {
        const data = await isSessionFoundBoolean();
        if (data===undefined) {
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
     
        }else{
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
            });

            // ajouter un commentair
            const formComment=document.querySelectorAll('.form-comment');
            addComment(formComment);

            // filter posts
            const contenCatId=document.querySelectorAll('.contenCatId');
            const categoryId=document.querySelectorAll('.categoryId');
            const contentPostBlock=document.querySelector('.content-post-block');

            filterPost(contenCatId,categoryId, function(_data) { 
                console.log(_data);
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
            })
         

        }
    }

    (async () => {
        await checkUserSession();
    })();
  


}); 

}

export{
    main
}