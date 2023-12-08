

const displayFormMecanisme=(formSignUp,formSignIn,row)=>{
    formSignUp.style.display="none";
    if (row[0]) {    
        row[0].addEventListener('click',()=>{
           
            if (formSignIn) {
                formSignIn.style.display="none";                            
                formSignUp.style.display="block";
            }
        })
    }
    if (row[1]) {   
        row[1].addEventListener('click',()=>{
            if (formSignUp) {
                formSignUp.style.display="none";
                formSignIn.style.display="block";  
            }
        })
    }
}

const readMore=(cardDescription,readMoreButton,onePostBlocks)=>{
    
    if (cardDescription && readMoreButton) {
           for (let j = 0; j < cardDescription.length; j++) {
           const desc = cardDescription[j];
           readMoreButton[j].addEventListener("click", function () {
            console.log(onePostBlocks[j]);
               desc.classList.toggle("expanded");
               readMoreButton[j].textContent = desc.classList.contains("expanded")
               ? "Read less"
               : "Read more";
               if (desc.classList.contains("expanded")) {
                onePostBlocks[j].style.height="500px";
                readMoreButton[j].textContent ="Read less";
               }else{
                readMoreButton[j].textContent ="Read more";
                onePostBlocks[j].style.height="400px";

               }
           });
           }
       }
}

const displayFomPost=(btn,modal,span)=>{

    btn.onclick = function() {
        modal.style.display = "block";
        }
    if (span) {           
        span.onclick = function() {
        modal.style.display = "none";
        }
    }
    
        window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
        }
}

const seeMore=(cardDescription,readMoreButton)=>{
  
    if (cardDescription && readMoreButton) {
        for (let j = 0; j < cardDescription.length; j++) {
            const desc = cardDescription[j];
            if (cardDescription[j].textContent.trim().length<100) {
                readMoreButton[j].style.display='none';
            }

        }
}
}


const disPlayComment=(comments,createCommentForm,lastPost,lastFormComment,lastBlocComment)=>{
    for (let k = 0; k < comments.length; k++) {
        comments[k].addEventListener("click", function () {
            // alert('yes')
            if (k!== comments.length-1 && createCommentForm[k].style.display==="none" ) {
                createCommentForm[k].style.display="block";              
            }else{
                if (k!== comments.length-1) createCommentForm[k].style.display="none" ;               
            }
            if (k===comments.length-1 && createCommentForm[k].style.display==="none"  && lastFormComment.style.display==="none" && lastBlocComment.children<2) {
                createCommentForm[k].style.display="block";
                lastPost.style.marginBottom="0px";
                lastFormComment.style.marginBottom="400px";
                return
            }else{
                if (k===comments.length-1&& createCommentForm[k].style.display==="block" ) {                                    
                    createCommentForm[k].style.display="none" 
                    lastPost.style.marginBottom="400px";
                    return
                }
            }
            if (k===comments.length-1) {
                createCommentForm[k].style.display="block";
                lastPost.style.marginBottom="0px";
                lastFormComment.style.marginBottom="0px";
                lastBlocComment.style.marginBottom="400px";
                return
            }else{
                if (k===comments.length-1 && createCommentForm[k].style.display==="block") {                    
                    lastBlocComment.style.marginBottom="0px";
                    lastPost.style.marginBottom="400px";
                }
            }
        });
    }
}

const disPlayCommentFilter=(comments,createCommentForm,blocComment,lastPost,lastFormComment,lastBlocComment)=>{
    for (let k = 0; k < comments.length; k++) {
        comments[k].addEventListener("click", function () {
            if (k!== comments.length-1 && createCommentForm[k].style.display==="none" && blocComment[k].style.display==="none"  ) {
                createCommentForm[k].style.display="block";    
                blocComment[k].style.display="block";            
            }else{
                if (k!== comments.length-1){
                     createCommentForm[k].style.display="none" ;  
                     blocComment[k].style.display="none" ;   
                }            
            }
            if (k===comments.length-1 && createCommentForm[k].style.display==="none" && blocComment[k].style.display==="none" ) {
                createCommentForm[k].style.display="block";
                blocComment[k].style.display="block"; 
                lastPost.style.marginBottom="0px";
                lastBlocComment.style.marginBottom="400px";
               
            }else{
                if (k===comments.length-1) {
                    createCommentForm[k].style.display="none";
                    blocComment[k].style.display="none"; 
                    lastPost.style.marginBottom="400px";
                    lastBlocComment.style.marginBottom="0px";
                }

            }

        });
    }
}

export{displayFormMecanisme,readMore,displayFomPost,seeMore,disPlayComment,disPlayCommentFilter}