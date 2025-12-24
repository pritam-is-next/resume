<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title><?= $$Title ?></title>
    <!-- Bootstrap 5.3 CSS -->
    <link rel="stylesheet" href="/css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/css/Bootstrap-Icons/bootstrap-icons.css">
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
                    <span class="fs-4 fw-bold"><i class="bi bi-person-circle me-2"></i><?= $$Heading ?></span>
                </a>
                <ul class="nav nav-pills flex-column mb-auto">
                </ul>
                <hr class="border-secondary">
                <div class="dropdown">
                    <a href="#" class="d-flex align-items-center text-body text-decoration-none dropdown-toggle"
                        data-bs-toggle="dropdown">
                        <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVA_HrQLjkHiJ2Ag5RGuwbFeDKRLfldnDasw&s" alt="Admin"
                            class="rounded-circle me-2" width="32">
                        <strong><?= $$User_Details->FullName ?></strong>
                    </a>
                    <ul class="dropdown-menu shadow">
                        <li><a class="dropdown-item" href="admin/profile"><i class="bi bi-person me-2"></i>Profile</a></li>
                        <li><a class="dropdown-item" href="#"><i class="bi bi-sliders me-2"></i>Settings</a></li>
                        <li>
                            <hr class="dropdown-divider">
                        </li>
                        <li>
                            <button class="dropdown-item" id="themeToggle">
                                <i class="bi bi-sun-fill me-2"></i>Light Mode
                            </button>
                        </li>
                        <li><a class="dropdown-item text-danger" href="/logout"><i
                                    class="bi bi-box-arrow-right me-2"></i>Logout</a></li>
                    </ul>
                </div>
            </div>
        </section>

        <!-- Main Content -->
        <div class="flex-grow-1 p-4 bg-gray-900">
            <div class="flex-grow-1 p-4 bg-gray-900">

    <form method="POST" action="/admin/profile" enctype="multipart/form-data">
        <div class="row g-4">

            <!-- Profile Info -->
            <div class="col-lg-8">
                <div class="card bg-body border-secondary shadow-sm">
                    <div class="card-header bg-body-tertiary border-secondary">
                        <h5 class="mb-0"><i class="bi bi-person-lines-fill me-2"></i>Profile Information</h5>
                    </div>
                    <div class="card-body">

                        <div class="row g-3">
                            <div class="col-md-6">
                                <label class="form-label">Full Name</label>
                                <input type="text" name="full_name" class="form-control"
                                       value="<?= $$User_Details->FullName ?>" required>
                            </div>

                            <div class="col-md-6">
                                <label class="form-label">Email</label>
                                <input type="email" name="email" class="form-control"
                                       value="<?= $$User_Details->Email ?>">
                            </div>

                            <div class="col-md-6">
                                <label class="form-label">Phone</label>
                                <input type="text" name="phone" class="form-control"
                                       value="<?= $$User_Details->Phone ?>">
                            </div>

                            <div class="col-md-6">
                                <label class="form-label">Date of Birth</label>
                                <input type="date" name="dob" class="form-control"
                                       value="<?= $$User_Details->Dob ?>">
                            </div>

                            <div class="col-md-6">
                                <label class="form-label">Gender</label>
                                <select name="gender" class="form-select">
                                    <option value="Male">Male</option>
                                    <option value="Female">Female</option>
                                    <option value="Other">Other</option>
                                </select>
                            </div>

                            <div class="col-12">
                                <label class="form-label">Bio</label>
                                <textarea name="bio" rows="4" class="form-control"><?= $$User_Details->Bio ?></textarea>
                            </div>
                        </div>

                    </div>
                </div>
            </div>

            <!-- Avatar & Address -->
            <div class="col-lg-4">
                <div class="card bg-body border-secondary shadow-sm mb-4">
                    <div class="card-header bg-body-tertiary border-secondary">
                        <h5 class="mb-0"><i class="bi bi-image me-2"></i>Avatar</h5>
                    </div>
                    <div class="card-body text-center">
                        <img src="<?= $$User_Details->Avatar ?>"
                             class="rounded-circle mb-3" width="120" height="120">
                        <input type="file" name="avatar" class="form-control">
                    </div>
                </div>

                <div class="card bg-body border-secondary shadow-sm">
                    <div class="card-header bg-body-tertiary border-secondary">
                        <h5 class="mb-0"><i class="bi bi-geo-alt me-2"></i>Address</h5>
                    </div>
                    <div class="card-body">

                        <div class="mb-2">
                            <label class="form-label">Address Line</label>
                            <input type="text" name="address_line" class="form-control"
                                   value="<?= $$User_Details->AddressLine ?>">
                        </div>

                        <div class="mb-2">
                            <label class="form-label">City</label>
                            <input type="text" name="city" class="form-control"
                                   value="<?= $$User_Details->City ?>">
                        </div>

                        <div class="mb-2">
                            <label class="form-label">State</label>
                            <input type="text" name="state" class="form-control"
                                   value="<?= $$User_Details->State ?>">
                        </div>

                        <div class="mb-2">
                            <label class="form-label">Country</label>
                            <input type="text" name="country" class="form-control"
                                   value="<?= $$User_Details->Country ?>">
                        </div>

                        <div class="mb-2">
                            <label class="form-label">Zip Code</label>
                            <input type="text" name="zip_code" class="form-control"
                                   value="<?= $$User_Details->ZipCode ?>">
                        </div>

                    </div>
                </div>
            </div>

            <!-- Save Button -->
            <div class="col-12">
                <div class="d-flex justify-content-end">
                    <button type="submit" class="btn btn-primary px-4">
                        <i class="bi bi-save me-1"></i>Save Changes
                    </button>
                </div>
            </div>

        </div>
    </form>

</div>

        </div>
    </div>

    <!-- Bootstrap JS -->
    <script src="/js/Bootstrap/bootstrap.bundle.min.js"></script>
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