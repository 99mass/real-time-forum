import {  routes,addRouter,replaceRouter,currentPath } from "../router/route.js";


const displayFormMecanisme=(formSignUp,formSignIn,row)=>{
    formSignUp.style.display="none";
    if (row[0]) {    
        row[0].addEventListener('click',()=>{
           
            if (formSignIn) {
                formSignIn.style.display="none";                            
                formSignUp.style.display="block";
                let r=routes["/Registre"]['name'];
                replaceRouter(r);
            }
        })
    }
    if (row[1]) {   
        row[1].addEventListener('click',()=>{
            if (formSignUp) {
                formSignUp.style.display="none";
                formSignIn.style.display="block";  
                let r=routes["/Login"]['name'];
                replaceRouter(r);
            }
        })
    }
}


const readMore=(cardDescription,readMoreButton,onePostBlocks)=>{
    
    if (cardDescription && readMoreButton) {
           for (let j = 0; j < cardDescription.length; j++) {
           const desc = cardDescription[j];
           desc.classList.add("hide-excess-lines"); 
           readMoreButton[j].addEventListener("click", function () {
               desc.classList.toggle("expanded");
               desc.classList.toggle("hide-excess-lines"); 
               readMoreButton[j].textContent = desc.classList.contains("expanded")
               ? "Read less"
               : "Read more";
               if (desc.classList.contains("expanded")) {
                onePostBlocks[j].style.height="fit-content";
                readMoreButton[j].textContent ="Read less";
               }else{
                readMoreButton[j].textContent ="Read more";
                onePostBlocks[j].style.height="fit-content";

               }
           });
           }
       }
}

const seeMore=(cardDescription,readMoreButton)=>{
    
    if (cardDescription && readMoreButton) {
        for (let j = 0; j < cardDescription.length; j++) {
            var lines=[];
            const desc = cardDescription[j];
             lines=desc.textContent.split(/\n|\r/);
            if (lines[0]=="") lines.shift();
            if (lines[lines.length]=="") lines.pop();
            const lineCount = lines.length;

            if (lineCount <=1) {
                readMoreButton[j].style.display='none';
            } else {
                readMoreButton[j].style.display='block';
            }
        }
       
    }
}
const reloadPost=()=>{
    document.querySelector('#all-categories').addEventListener('click',()=>{
        window.location.reload();
    })
}

const displayFomPost=(btn,modal,span)=>{
    let r1=routes["/Home"]['name'];
    let r2=routes["/AddPost"]['name'];
    
   if (currentPath==="/AddPost") {
      modal.style.display = "block";
      addRouter(r2);      
   }
   
    btn.onclick = function() {
        modal.style.display = "block";
        replaceRouter(r2);
        }
    if (span) {           
        span.onclick = function() {
        modal.style.display = "none";
        replaceRouter(r1);

        }
    }
    
        window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
            replaceRouter(r1);
        }
        }
}




const disPlayComment=(comments,createCommentForm)=>{
    for (let k = 0; k < comments.length; k++) {
        comments[k].addEventListener("click", function () {           
            if ( createCommentForm[k].style.display==="none" ) {
                createCommentForm[k].style.display="block"; 
            }else{
                createCommentForm[k].style.display="none"; 
            }
        });
    }
}

const disPlayCommentFilter=(comments,createCommentForm,blocComment)=>{
    for (let k = 0; k < comments.length; k++) {
        comments[k].addEventListener("click", function () {
            if ( createCommentForm[k].style.display==="none" ) {
                createCommentForm[k].style.display="block"; 
                blocComment[k].style.display="block";   
            }else{
                createCommentForm[k].style.display="none" ;
                blocComment[k].style.display="none" ;   
            }

        });
    }
}

export{displayFormMecanisme,readMore,displayFomPost,seeMore,disPlayComment,disPlayCommentFilter,reloadPost}