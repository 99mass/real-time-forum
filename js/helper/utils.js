const isGoodNumber=(score)=>{
    let likecount = Number(score);
    if (isNaN(likecount)) {
        console.error("Error: The value must be a number");
        return null
      } else if (!Number.isInteger(likecount) || likecount<0) {
        console.error("Error: The value must be an integer");
        return null
    } 
     return likecount;

}


export{isGoodNumber}