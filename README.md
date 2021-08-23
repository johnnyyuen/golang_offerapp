# golang_offerapp

## Dev start up file
you can create a startup file for your environment (e.g. dev)

```
#!/bin/bash

export TOKEN_SECRET="secretfortoken"
export DBNAME=offerapp
export DBUSER=postgres
export DBPASSWORD=postgres

echo "starting go server at port 3000"

go run main.go
```

and then make sure it is executable by chmod 755 dev


Issue and trouble shooting

Got error about:
> Error illegal base64 data at input byte 0

Fixed by "Bearer " instead of "Bearer" in AuthMiddleWare function

```
<script>
  

async function handleFormSubmit(event) {
	event.preventDefault();

	const form = event.currentTarget;
	const url = form.action;

	try {
		const formData = new FormData(form);
		const responseData = await postFormDataAsJson({ url, formData });

		console.log({ responseData });
	} catch (error) {
		console.error(error);
	}
}

async function postFormDataAsJson({url, formData }) {
  const plainFormData = Object.fromEntries(formData.entries());
	const formDataJsonString = JSON.stringify(plainFormData);

  const fetchOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    body: formDataJsonString,
  };
  const response = await fetch(url, fetchOptions);

  if (!response.ok) {
    const errorMessage = await response.text();
    throw new Error(errorMessage);
  }

return response.json();
}




  const form = document.getElementById("login-form");
  form.addEventListener('submit', handleFormSubmit);
</script>
```