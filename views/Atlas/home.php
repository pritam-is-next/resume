<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">
<head>
    <meta charset="UTF-8">
    <title><?= $$Title ?> — Atlas</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="stylesheet" href="/css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/css/Bootstrap-Icons/bootstrap-icons.css">

    <style>
        body {
            background: #0f1220;
            color: #e5e7eb;
            font-family: Inter, system-ui, sans-serif;
        }

        section {
            padding: 5rem 0;
        }

        h1, h2 {
            font-weight: 700;
            letter-spacing: -0.02em;
        }

        .hero {
            background: radial-gradient(circle at top right, #4f46e5, transparent 60%);
        }

        .card {
            background: #161a2f;
            border: 1px solid #232857;
            border-radius: 20px;
        }

        .badge-skill {
            background: #232857;
            border-radius: 999px;
            padding: .5rem 1rem;
            margin: .25rem;
            display: inline-block;
            font-size: .85rem;
        }

        .project-card {
            overflow: hidden;
        }

        .project-card img {
            width: 100%;
            height: 180px;
            object-fit: cover;
        }

        .cta {
            background: linear-gradient(135deg, #4f46e5, #06b6d4);
            border-radius: 24px;
            padding: 4rem 2rem;
            text-align: center;
        }

        footer {
            color: #9ca3af;
            text-align: center;
            padding: 2rem 0;
        }
    </style>
</head>

<body>

<main>

    <!-- HERO -->
    <section class="hero">
        <div class="container">
            <div class="row align-items-center g-5">
                <div class="col-md-6">
                    <h1 class="display-4"><?= $$Hero->Heading ?></h1>
                    <p class="lead text-muted"><?= $$Hero->SubHeading ?></p>

                    <?php foreach ($$Hero->CallToActions as $cta): ?>
                        <a href="<?= $cta->Href ?>" class="btn btn-primary btn-lg me-2">
                            <?= $cta->Text ?>
                        </a>
                    <?php endforeach; ?>
                </div>

                <div class="col-md-6 text-center">
                    <img src="<?= $$Avatar ?>" class="img-fluid rounded-4 shadow">
                </div>
            </div>
        </div>
    </section>

    <!-- ABOUT -->
    <section>
        <div class="container">
            <div class="card p-5">
                <h2>About Me</h2>
                <p class="lead"><?= $$AboutMe ?></p>
                <p class="text-muted">
                    📍 <?= $$ContactDetails->Location ?> ·
                    ✉ <?= $$ContactDetails->Email ?> ·
                    ☎ <?= $$ContactDetails->Phone ?>
                </p>
            </div>
        </div>
    </section>

    <!-- SKILLS -->
    <section>
        <div class="container">
            <h2 class="mb-4">Skills</h2>
            <?php foreach ($$Skills as $skill): ?>
                <span class="badge-skill">
                    <?= $skill->Name ?> · <?= $skill->Level ?>%
                </span>
            <?php endforeach; ?>
        </div>
    </section>

    <!-- EXPERIENCE -->
    <section>
        <div class="container">
            <h2 class="mb-5">Experience</h2>

            <?php foreach ($$Experiences as $exp): ?>
                <div class="card p-4 mb-4">
                    <h5><?= $exp->Designation ?></h5>
                    <small class="text-info"><?= $exp->Duration ?></small>
                    <p class="mt-3"><?= $exp->Description ?></p>
                </div>
            <?php endforeach; ?>
        </div>
    </section>

    <!-- PROJECTS -->
    <section>
        <div class="container">
            <h2 class="mb-5">Projects</h2>
            <div class="row g-4">

                <?php foreach ($$Projects as $project): ?>
                    <div class="col-md-6 col-lg-4">
                        <div class="card project-card h-100">
                            <img src="<?= $project->Image ?>" alt="">
                            <div class="p-4">
                                <h5><?= $project->Title ?></h5>
                                <p class="text-muted"><?= $project->SmallDescription ?></p>
                                <a href="<?= $project->Link ?>" class="btn btn-outline-primary btn-sm">
                                    View Project
                                </a>
                            </div>
                        </div>
                    </div>
                <?php endforeach; ?>

            </div>
        </div>
    </section>

    <!-- CONTACT CTA -->
    <section>
        <div class="container">
            <div class="cta">
                <h2>Let’s Work Together</h2>
                <p class="lead">
                    Reach out via email or connect on GitHub & LinkedIn.
                </p>
                <a href="mailto:<?= $$ContactDetails->Email ?>" class="btn btn-dark btn-lg">
                    Contact Me
                </a>
            </div>
        </div>
    </section>

</main>

<footer>
    © 2026 · <?= $$Title ?> · Atlas Theme
</footer>

</body>
</html>
