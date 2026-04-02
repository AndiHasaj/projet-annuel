<?php
session_start();

// Langue par défaut
$lang = $_SESSION['lang'] ?? 'fr';

// Si l'utilisateur change de langue via l'URL (?lang=en)
if (isset($_GET['lang'])) {
    $lang = $_GET['lang'];
    $_SESSION['lang'] = $lang;
}

// Chargement du fichier JSON
$json_file = file_get_contents(__DIR__ . "/../langues/{$lang}.json");
$textes = json_decode($json_file, true);
?>