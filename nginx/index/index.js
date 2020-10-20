function subscribe() {
	const url = document.getElementById("url").value;
  const email = document.getElementById("email").value;
  let form = {
  	"url": url,
    "email": email
  };
  
  var xhr = new XMLHttpRequest();
  
  xhr.open("POST", '/submit', true);
  xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
  xhr.send(JSON.stringify(form));
}