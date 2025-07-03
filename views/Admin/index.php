<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>CMS Admin Panel</title>
    <!-- Bootstrap 5.3 CSS -->
    <link rel="stylesheet" href="/Css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/Css/Bootstrap-Icons/bootstrap-icons.css">
    <style>
        /* Custom scrollbar for dark mode */
        ::-webkit-scrollbar {
            width: 8px;
        }

        ::-webkit-scrollbar-track {
            background: #212529;
        }

        ::-webkit-scrollbar-thumb {
            background: #495057;
            border-radius: 4px;
        }
    </style>
</head>

<body class="bg-body">
    <!-- Wrapper -->
    <div class="d-flex">
        <!-- Offcanvas Sidebar -->
        <section id="sidebar"
            class="offcanvas-lg offcanvas-start bg-body border-end border-secondary min-vh-100"
            tabindex="-1" style="width: 280px;">
            <div class="offcanvas-body d-flex flex-column h-100 p-3">
                <a href="" class="d-flex align-items-center mb-3 text-body text-decoration-none">
                    <span class="fs-4 fw-bold"><i class="bi bi-database-fill-gear me-2"></i>CMS Admin</span>
                </a>
                <hr class="border-secondary">
                <ul class="nav nav-pills flex-column mb-auto">
                    <li class="nav-item">
                        <a href="#" class="nav-link active bg-body-secondary text-body" aria-current="page">
                            <i class="bi bi-speedometer2 me-2"></i>Dashboard
                        </a>
                    </li>
                    <li><a href="#" class="nav-link text-body"><i class="bi bi-file-earmark-text me-2"></i>Posts</a>
                    </li>
                    <li><a href="#" class="nav-link text-body"><i class="bi bi-people me-2"></i>Users</a></li>
                    <li><a href="#" class="nav-link text-body"><i class="bi bi-images me-2"></i>Media</a></li>
                    <li><a href="#" class="nav-link text-body"><i class="bi bi-gear me-2"></i>Settings</a></li>
                </ul>
                <hr class="border-secondary">
                <div class="dropdown">
                    <a href="#" class="d-flex align-items-center text-body text-decoration-none dropdown-toggle"
                        data-bs-toggle="dropdown">
                        <img src="https://via.placeholder.com/32/495057/ffffff?text=AD" alt="Admin"
                            class="rounded-circle me-2" width="32">
                        <strong>Admin</strong>
                    </a>
                    <ul class="dropdown-menu shadow">
                        <li><a class="dropdown-item" href="#"><i class="bi bi-person me-2"></i>Profile</a></li>
                        <li><a class="dropdown-item" href="#"><i class="bi bi-sliders me-2"></i>Settings</a></li>
                        <li>
                            <hr class="dropdown-divider">
                        </li>
                        <li>
                            <button class="dropdown-item" id="themeToggle">
                                <i class="bi bi-sun-fill me-2"></i>Light Mode
                            </button>
                        </li>
                        <li><a class="dropdown-item text-danger" href="#"><i
                                    class="bi bi-box-arrow-right me-2"></i>Logout</a></li>
                    </ul>
                </div>
            </div>
        </section>

        <!-- Main Content -->
        <div class="flex-grow-1 p-4 bg-gray-900">
            <!-- Top Bar -->
            <div class="d-flex justify-content-between align-items-center mb-4">
                <h2><i class="bi bi-speedometer2 me-2"></i>Dashboard</h2>
                <div>
                    <button class="btn btn-sm btn-outline-light me-2 d-lg-none" data-bs-toggle="offcanvas"
                        data-bs-target="#sidebar">
                        <i class="bi bi-list"></i>
                    </button>
                    <button class="btn btn-sm btn-primary">
                        <i class="bi bi-plus-lg me-1"></i> New Post
                    </button>
                </div>
            </div>

            <!-- Stats Cards -->
            <div class="row g-4 mb-4">
                <div class="col-md-3">
                    <div class="card bg-body border-primary">
                        <div class="card-body">
                            <div class="d-flex justify-content-between align-items-center">
                                <div>
                                    <h6 class="card-title text-primary">Total Posts</h6>
                                    <p class="fs-2 mb-0">128</p>
                                </div>
                                <i class="bi bi-file-earmark-text fs-1 text-primary opacity-25"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="card bg-body border-success">
                        <div class="card-body">
                            <div class="d-flex justify-content-between align-items-center">
                                <div>
                                    <h6 class="card-title text-success">Active Users</h6>
                                    <p class="fs-2 mb-0">42</p>
                                </div>
                                <i class="bi bi-people fs-1 text-success opacity-25"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="card bg-body border-warning">
                        <div class="card-body">
                            <div class="d-flex justify-content-between align-items-center">
                                <div>
                                    <h6 class="card-title text-warning">Media Files</h6>
                                    <p class="fs-2 mb-0">316</p>
                                </div>
                                <i class="bi bi-images fs-1 text-warning opacity-25"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="card bg-body border-info">
                        <div class="card-body">
                            <div class="d-flex justify-content-between align-items-center">
                                <div>
                                    <h6 class="card-title text-info">Comments</h6>
                                    <p class="fs-2 mb-0">87</p>
                                </div>
                                <i class="bi bi-chat-left-text fs-1 text-info opacity-25"></i>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Recent Activity -->
            <div class="row g-4">
                <!-- Recent Posts Table -->
                <div class="col-lg-8">
                    <div class="card bg-body border-secondary shadow-sm">
                        <div class="card-header bg-body-tertiary border-secondary">
                            <h5 class="mb-0"><i class="bi bi-file-earmark-text me-2"></i>Recent Posts</h5>
                        </div>
                        <div class="card-body">
                            <div class="table-responsive">
                                <table class="table bg-primary table-hover align-middle">
                                    <thead class="bg-gray-800">
                                        <tr>
                                            <th>ID</th>
                                            <th>Title</th>
                                            <th>Status</th>
                                            <th>Date</th>
                                            <th>Actions</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>1</td>
                                            <td>Getting Started with CMS</td>
                                            <td><span class="badge bg-success">Published</span></td>
                                            <td>2023-10-15</td>
                                            <td>
                                                <button class="btn btn-sm btn-outline-primary me-1"><i
                                                        class="bi bi-pencil-square"></i></button>
                                                <button class="btn btn-sm btn-outline-danger"><i
                                                        class="bi bi-trash"></i></button>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td>2</td>
                                            <td>Advanced Content Types</td>
                                            <td><span class="badge bg-warning text-dark">Draft</span></td>
                                            <td>2023-10-14</td>
                                            <td>
                                                <button class="btn btn-sm btn-outline-primary me-1"><i
                                                        class="bi bi-pencil-square"></i></button>
                                                <button class="btn btn-sm btn-outline-danger"><i
                                                        class="bi bi-trash"></i></button>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Quick Actions Widget -->
                <div class="col-lg-4">
                    <div class="card bg-body border-secondary shadow-sm">
                        <div class="card-header bg-gray-800 border-secondary">
                            <h5 class="mb-0"><i class="bi bi-lightning-charge me-2"></i>Quick Actions</h5>
                        </div>
                        <div class="card-body">
                            <button class="btn btn-outline-primary w-100 mb-2">
                                <i class="bi bi-file-earmark-plus me-2"></i>New Post
                            </button>
                            <button class="btn btn-outline-success w-100 mb-2">
                                <i class="bi bi-upload me-2"></i>Upload Media
                            </button>
                            <button class="btn btn-outline-info w-100 mb-2">
                                <i class="bi bi-person-plus me-2"></i>Add User
                            </button>
                            <hr class="border-secondary">
                            <div class="alert alert-dark border-secondary">
                                <i class="bi bi-info-circle me-2"></i>
                                <strong>Tip:</strong> Press <kbd>Ctrl + N</kbd> to create new content
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Bootstrap JS -->
    <script src="/Js/Bootstrap/bootstrap.bundle.min.js"></script>
    <!-- Theme Toggle Script -->
    <script>
        document.getElementById('themeToggle').addEventListener('click', function() {
            const htmlEl = document.documentElement;
            const currentTheme = htmlEl.getAttribute('data-bs-theme');
            const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
            htmlEl.setAttribute('data-bs-theme', newTheme);
            this.innerHTML = newTheme === 'dark' ?
                '<i class="bi bi-sun-fill me-2"></i>Light Mode' :
                '<i class="bi bi-moon-stars-fill me-2"></i>Dark Mode';
        });
    </script>
</body>

</html>