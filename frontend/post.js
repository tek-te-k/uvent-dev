const signup = async () => {
  name = document.getElementById("name").value;
  email = document.getElementById("email").value;
  birth = document.getElementById("birth").value;
  address = document.getElementById("address").value;
  password = document.getElementById("password").value;
  password_confirmation = document.getElementById("passwordConfirm").value;
  alert(birth);

  if (password !== password_confirmation) {
    alert("Passwords do not match");
    return;
  }
  let reqBody = JSON.stringify({
    name,
    email,
    birth,
    address,
    password,
    password_confirmation,
  });
  try {
    const response = await fetch("http://localhost:8080/auth/signup", {
      method: "POST",
      body: reqBody,
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
    });
  } catch (e) {
    alert("bad request");
    console.log(e);
  }
};
