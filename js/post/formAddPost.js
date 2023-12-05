

const formAddPost=(formCreatPost)=>{
    formCreatPost.addEventListener('submit', function(e) {
        e.preventDefault();

        let selectedCategories = [];
        document.querySelectorAll('input[name="categories"]:checked').forEach((checkbox) => {

            selectedCategories.push(checkbox.value);
        });

        let title = document.querySelector('input[name="Title"]').value.trim();
        let content = document.querySelector('textarea[name="Content"]').value.trim();
        let file = document.querySelector('input[type="file"]').files[0];

        if (file) {
            let fileType = file.type;
            let fileSize = file.size / (1024 * 1024); // size in MB

            if (!(['image/jpeg', 'image/svg+xml', 'image/png', 'image/gif'].includes(fileType) && fileSize <= 20)) {
                alert('File is not valid')
                return;
            }
        }
        console.log(selectedCategories);
        let reader = new FileReader();
        reader.onloadend = function() {
            let base64File = reader.result;
            let data = {
                Title: title,
                Content: content,
                Category: selectedCategories 
            };
        
            if (file) {
                data.image = base64File;
            }
        
            fetch('http://localhost:8080/addpost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response =>{ 
                if (response.status===200) {
                    window.location.reload();
                }
                return  response.json()
            })
            .then(data =>{ 
                if (data["message"]) {
                    
                    alert(data["message"]);
                }
                console.log(data)
            })
            .catch((error) => {
                console.error('Error:', error);
            });
        }
        
        if (file) {
            reader.readAsDataURL(file);
        } else {
            reader.onloadend();
        }

        console.log('Selected Categories:', selectedCategories);
        console.log('Title:', title);
        console.log('Content:', content);
        console.log('File:', file);
    });


}

export{formAddPost}



// const  errorPost=document.querySelector('.error-post');
// errorPost.style.diplay="block";
// errorPost.innerHTML='Content is required';