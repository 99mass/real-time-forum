import { linkApi } from "../../helper/api_link.js";
import { isGoodNumber } from "../../helper/utils.js";


// const liskePost=(likePost,likePostId,likePostScore,dislikePostScore,dislikePost)=>{
   
//     for (let i = 0; i < likePost.length; i++) {
//         const btnLikePost = likePost[i];
//        btnLikePost.addEventListener('click', async function() {
        
//         let postID=likePostId[i].textContent.trim();
//             try {
//                 const response = await fetch(`${linkApi}likepost`, {
//                     method: 'POST',
//                     headers: {
//                         'Content-Type': 'application/json',
//                     },
//                     body: JSON.stringify({PostID: postID}),

//                 });

//                 if(response.status === 200){

//                    let countLike=isGoodNumber(likePostScore[i].textContent.trim());
//                    let countDisLike=isGoodNumber(dislikePostScore[i].textContent.trim());
//                     if (countDisLike>0) countDisLike-=1;
//                    if (countLike===null) return ;
                   
//                     if (!btnLikePost.className.includes('like-post-color')) {                    
//                         likePostScore[i].textContent=countLike+1;
//                         btnLikePost.style.color="darkgoldenrod";
//                         likePostScore[i].style.color="darkgoldenrod";

//                         dislikePostScore[i].textContent=countDisLike;
//                         dislikePostScore[i].style.color="#ffffff";
//                     }

//                 }else{

//                     const er = await response.json();
//                     console.log('Error : ' + er.message);
//                     return;
//                 }
                
//             } catch (error) {
//             console.log("error : "+error);
//             }
//         });
        
//     }
// }


const liskePost=(likePost,likePostId,likePostScore,dislikePostScore,dislikePost)=>{
   
    for (let i = 0; i < likePost.length; i++) {
        const btnLikePost = likePost[i];

        const handleClick = async function() {
            alert('ok')
            let postID=likePostId[i].textContent.trim();
            try {
                const response = await fetch(`${linkApi}likepost`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({PostID: postID}),

                });

                if(response.status === 200){

                   let countLike=isGoodNumber(likePostScore[i].textContent.trim());
                   let countDisLike=isGoodNumber(dislikePostScore[i].textContent.trim());
                    if (countDisLike>0) countDisLike-=1;
                   if (countLike===null) return ;
                   
                    if (!btnLikePost.className.includes('like-post-color')) {                    
                        likePostScore[i].textContent=countLike+1;
                        btnLikePost.style.color="darkgoldenrod";
                        likePostScore[i].style.color="darkgoldenrod";

                        dislikePostScore[i].textContent=countDisLike;
                        dislikePost[i].style.color="#ffffff";
                        dislikePostScore[i].style.color="#ffffff";

                        btnLikePost.classList.add('like-post-color');
                        dislikePost[i].classList.remove('dislike-post-color');
                    }

                }else{

                    const er = await response.json();
                    console.log('Error : ' + er.message);
                    return;
                }
                
            } catch (error) {
            console.log("error : "+error);
            }
        };

        btnLikePost.addEventListener('click', handleClick);
        
    }
}


export{
    liskePost
}