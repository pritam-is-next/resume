<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">
<head>
    <meta charset="UTF-8">
    <title><?= $$Title ?> — CV</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="stylesheet" href="/css/Bootstrap/bootstrap.min.css">
    <link rel="stylesheet" href="/css/Bootstrap-Icons/bootstrap-icons.css">

    <style>
        body {
            background: #0d1117;
            color: #c9d1d9;
            font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
        }

        a { color: #58a6ff; text-decoration: none; }
        a:hover { text-decoration: underline; }

        .container-narrow {
            max-width: 900px;
            margin: auto;
            padding: 4rem 1.5rem;
        }

        h1, h2 {
            color: #f0f6fc;
            border-bottom: 1px solid #30363d;
            padding-bottom: .5rem;
            margin-bottom: 1.5rem;
        }

        section {
            margin-bottom: 4rem;
        }

        .meta {
            color: #8b949e;
            font-size: .9rem;
        }

        .skill {
            display: flex;
            justify-content: space-between;
            border-bottom: 1px dashed #30363d;
            padding: .4rem 0;
        }

        .project {
            margin-bottom: 1.5rem;
        }

        footer {
            text-align: center;
            color: #8b949e;
            font-size: .85rem;
            padding-bottom: 3rem;
        }
    </style>
</head>

<body>

<main class="container-narrow">

    <!-- INTRO -->
    <section>
        <h1><?= $$Hero->Heading ?></h1>
        <p class="meta"><?= $$Hero->SubHeading ?></p>
    </section>

    <!-- ABOUT -->
    <section>
        <h2>About</h2>
        <p><?= $$AboutMe ?></p>
        <p class="meta">
            📍 <?= $$ContactDetails->Location ?> ·
            📧 <?= $$ContactDetails->Email ?> ·
            ☎ <?= $$ContactDetails->Phone ?>
        </p>
    </section>

    <!-- SKILLS -->
    <section>
        <h2>Skills</h2>
        <?php foreach ($$Skills as $skill): ?>
            <div class="skill">
                <span><?= $skill->Name ?></span>
                <span class="meta"><?= $skill->Level ?>%</span>
            </div>
        <?php endforeach; ?>
    </section>

    <!-- EXPERIENCE -->
    <section>
        <h2>Experience</h2>
        <?php foreach ($$Experiences as $Experience): ?>
            <div class="mb-4">
                <strong><?= $Experience->Designation ?></strong><br>
                <span class="meta"><?= $Experience->Duration ?></span>
                <p class="mt-2"><?= $Experience->Description ?></p>
            </div>
        <?php endforeach; ?>
    </section>

    <!-- PROJECTS -->
    <section>
        <h2>Projects</h2>
        <?php foreach ($$Projects as $Project): ?>
            <div class="project">
                <strong><?= $Project->Title ?></strong><br>
                <span class="meta"><?= $Project->SmallDescription ?></span><br>
                <a href="<?= $Project->Link ?>" target="_blank">View Source →</a>
            </div>
        <?php endforeach; ?>
    </section>

    <!-- CONTACT -->
    <section>
        <h2>Contact</h2>
        <p>
            Email: <a href="mailto:<?= $$ContactDetails->Email ?>"><?= $$ContactDetails->Email ?></a><br>
            GitHub: <a href="<?= $$ContactDetails->Github ?>"><?= $$ContactDetails->Github ?></a><br>
            LinkedIn: <a href="<?= $$ContactDetails->LinkedIn ?>"><?= $$ContactDetails->LinkedIn ?></a>
        </p>
    </section>

</main>

<footer>
    © <?= date("Y") ?> · <?= $$Title ?>
</footer>

</body>
</html>
