<!DOCTYPE html>
<html data-bs-theme="dark" lang="en">

<head>
  <title>OP-Resume</title>
  <link href="/css/Bootstrap/bootstrap.min.css" rel="stylesheet">
  <link href="/css/main.css" rel="stylesheet">
</head>

<body>
  <main>
    <section class="vh-100 bg-body">
      <div class="container h-100">
        <div class="row align-items-center h-100">
          <!-- Left column with text -->
          <div class="col-lg-6 mb-5 mb-lg-0 text-center text-lg-start">
            <h1 class="my-5 display-4 fw-bold text-primary">
              Welcome to OP Resume<br />
              <span class="text-secondary fs-1">Website Resume for Professionals</span>
            </h1>
            <p class="text-warning">
            It is a open soruce solution and everyone has the permissio to deploy this application in their solution and servers & until and unless they are earning money out of it    
          </p>
          </div>

          <!-- Right column with login form -->
          <div class="col-lg-6 d-flex justify-content-center">
            <div class="card shadow-sm w-100" style="max-width: 400px;">
              <div class="card-body p-5 bg-body-tertiary">
                <h4 class="mb-4 text-center">Register</h4>
                <form method="post">
                  <!-- Email input -->
                  <div class="form-floating mb-3">
                    <input type="email" class="form-control" id="loginEmail"
                      placeholder="name" name="loginEmail">
                    <label for="loginEmail">Email Address</label>
                  </div>

                  <!-- Password input -->
                  <div class="form-floating mb-4">
                    <input type="password" class="form-control" id="loginPassword"
                      placeholder="Password" name="loginPassword">
                    <label for="loginPassword">Password</label>
                  </div>

                  <!-- Password input -->
                  <div class="form-floating mb-4">
                    <input type="text" class="form-control" id="fullName"
                      placeholder="Full Name" name="loginPassword">
                    <label for="fullName">Full Name</label>
                  </div>

                  <!-- Submit button -->
                  <div class="d-grid mb-3">
                    <button type="submit" class="btn btn-primary">Login</button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

  </main>
  <footer>
    <script src="/Js/Bootstrap/bootstrap.min.js"></script>
  </footer>
</body>

</html>