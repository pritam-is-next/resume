<!doctype html>
<html lang="en" data-bs-theme="dark">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Admin Panel</title>
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>

<!-- Mobile Toggle Button -->
<button class="btn btn-primary position-fixed top-0 start-0 m-3 d-md-none z-3" type="button" data-bs-toggle="offcanvas" data-bs-target="#sidebar" aria-controls="sidebar" aria-label="Toggle navigation">
  ☰
</button>

<div class="container-fluid">
  <div class="row">
    <!-- Sidebar -->
    <nav id="sidebar" class="col-md-3 col-lg-2 d-md-block bg-dark text-white offcanvas-md offcanvas-start" tabindex="-1" aria-labelledby="sidebarLabel">
      <div class="offcanvas-header d-md-none">
        <h5 class="offcanvas-title" id="sidebarLabel">Menu</h5>
        <button type="button" class="btn-close btn-close-white" data-bs-dismiss="offcanvas" aria-label="Close"></button>
      </div>
      <div class="offcanvas-body p-3 p-md-0 d-flex flex-column">
        <ul class="nav nav-pills flex-column mb-auto">
          <li class="nav-item">
            <a href="#" class="nav-link active text-white">
              Dashboard
            </a>
          </li>
          <li>
            <a href="#" class="nav-link text-white">
              Posts
            </a>
          </li>
          <li>
            <a href="#" class="nav-link text-white">
              Media
            </a>
          </li>
          <li>
            <a href="#" class="nav-link text-white">
              Settings
            </a>
          </li>
        </ul>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4 py-4">
      <h1>Dashboard</h1>
      <p class="lead">Welcome to your admin panel.</p>
      <!-- Add more content here -->
    </main>
  </div>
</div>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
