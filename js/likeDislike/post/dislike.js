import { linkApi } from "../../helper/api_link.js";
import { isGoodNumber } from "../../helper/utils.js";


// const DisLiskePost=(dislikePost,dislikePostId,likePostScore,dislikePostScore)=>{
//     // console.log(dislikePost);
//     for (let i = 0; i < dislikePost.length; i++) {
//         const btnDisLikePost = dislikePost[i];
//         btnDisLikePost.addEventListener('click', async function() {
//                 let postID=dislikePostId[i].textContent.trim();
//                 try {
//                     const response = await fetch(`${linkApi}dislikepost`, {
//                         method: 'POST',
//                         headers: {
//                             'Content-Type': 'application/json',
//                         },
//                         body: JSON.stringify({PostID: postID}),

//                     });

//                     if(response.status === 200){
//                         let countDisLike =isGoodNumber( dislikePostScore[i].textContent.trim());
//                         let countLike =isGoodNumber(likePostScore[i].textContent.trim());
//                          if (countLike>0) countLike-=1;
//                         if (countDisLike===null) return ;
                        
//                          if (!btnDisLikePost.className.includes('like-post-color')) {                    
//                             dislikePostScore[i].textContent=countDisLike+1;
//                             btnDisLikePost.style.color="gray";
//                             dislikePostScore[i].style.color="gray";

//                              likePostScore[i].textContent=countDisLike;
//                              likePostScore[i].style.color="#ffffff";
//                          }

//                     }else{
//                         const er = await response.json();
//                         console.log('Error : ' + er.message);
//                         return;
//                     }
                    
//                 } catch (error) {
//                 console.log("error : "+error);
//                 }
//          });
//     }
//   }
  

const DisLiskePost=(dislikePost,dislikePostId,likePostScore,dislikePostScore,likePost)=>{
   
    for (let i = 0; i < dislikePost.length; i++) {
        const btnDisLikePost = dislikePost[i];

        const handleClick = async function() {
          alert('ok ok')
            let postID=dislikePostId[i].textContent.trim();
            try {
                const response = await fetch(`${linkApi}dislikepost`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({PostID: postID}),

                });

                if(response.status === 200){

                   let countDisLike =isGoodNumber( dislikePostScore[i].textContent.trim());
                   let countLike =isGoodNumber(likePostScore[i].textContent.trim());
                    if (countLike>0) countLike-=1;
                   if (countDisLike===null) return ;
                   
                    if (!btnDisLikePost.className.includes('dislike-post-color')) {                    
                        dislikePostScore[i].textContent=countDisLike+1;
                        btnDisLikePost.style.color="gray";
                        dislikePostScore[i].style.color="gray";

                        likePostScore[i].textContent=countLike;
                        likePostScore[i].style.color="#ffffff";
                        likePost[i].style.color="#ffffff";
                        btnDisLikePost.classList.add('dislike-post-color');
                        likePost[i].classList.remove('like-post-color');
                    }

                    return
                }else{

                    const er = await response.json();
                    console.log('Error : ' + er.message);
                    return;
                }
                
            } catch (error) {
            console.log("error : "+error);
            }
        };

        btnDisLikePost.addEventListener('click', handleClick);
        
    }
}

  export{DisLiskePost
  }