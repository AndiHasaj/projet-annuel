<?php
$id = 3; // normalement = session utilisateur en récuperant l'email et en faisant une requete pour récuperer l'id
    
$url = "http://localhost:8080/particuliers/" . $id;

$data = file_get_contents($url);
$particulier = json_decode($data, true);

if (!$particulier) {
    $particulier = [
        "nom" => "Inconnu",
        "score" => 0
    ];
}
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Tableau de Bord</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">

    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        'ucu-foret': '#1E4E4E',
                        'ucu-lagon': '#6A9A9A',
                        'ucu-grey': '#F8FAFC',
                    },
                    fontFamily: {
                        'sans': ['Inter', 'sans-serif'],
                        'montserrat': ['Montserrat', 'sans-serif'],
                    },
                    borderRadius: { 'ucu': '8px' }
                }
            }
        }
        
        function closeTuto() {
            document.getElementById('tuto-overlay').classList.add('hidden');
        }
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">

    <div id="tuto-overlay" class="fixed inset-0 z-[100] bg-ucu-foret/95 backdrop-blur-md flex items-center justify-center p-4">
        <div class="bg-white max-w-md w-full rounded-ucu shadow-2xl overflow-hidden">
            <div class="p-8 text-center">
                <div class="w-20 h-20 bg-ucu-lagon/10 rounded-full flex items-center justify-center mx-auto mb-6">
                    <span class="text-3xl text-ucu-lagon">♻️</span>
                </div>
                <h2 class="font-montserrat font-bold text-2xl text-ucu-foret uppercase mb-4 tracking-tighter">Bienvenue chez <br><span class="text-ucu-lagon">UpcycleConnect</span></h2>
                <p class="text-xs text-gray-500 leading-relaxed mb-8 italic">
                    Ici, vous pouvez donner, vendre vos objets et apprendre à transformer vos déchets en ressources.
                </p>
                <div class="space-y-3 mb-8">
                    <div class="flex items-center gap-3 text-left p-3 bg-ucu-grey rounded-ucu">
                        <span class="text-ucu-lagon font-bold italic">01.</span>
                        <p class="text-[10px] font-bold text-ucu-foret uppercase">Déposez une annonce ou un objet</p>
                    </div>
                    <div class="flex items-center gap-3 text-left p-3 bg-ucu-grey rounded-ucu">
                        <span class="text-ucu-lagon font-bold italic">02.</span>
                        <p class="text-[10px] font-bold text-ucu-foret uppercase">Gérez votre planning et vos cours</p>
                    </div>
                </div>
                <button onclick="closeTuto()" class="w-full bg-ucu-foret text-white py-4 rounded-ucu font-montserrat font-bold text-xs uppercase tracking-[0.2em] hover:bg-ucu-lagon transition-all">
                    Découvrir mon espace
                </button>
            </div>
        </div>
    </div>

    <nav class="bg-white border-b border-gray-100 sticky top-0 z-50">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold italic">UCU</div>
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest hidden sm:block">Particulier</h1>
            </div>
            <div class="flex items-center gap-8">
                <div class="text-right hidden md:block">
                    <p class="text-[9px] font-bold text-ucu-lagon uppercase">Mon Impact</p>
                    <p class="text-xs font-montserrat font-bold text-ucu-foret tracking-tighter italic">Score : <?= $particulier["score"] ?? 0 ?> pts</p>
                </div>
                <a href="particulier_profil.php" class="w-10 h-10 rounded-full bg-ucu-grey border border-gray-200 flex items-center justify-center cursor-pointer">👤</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
            
            <div class="lg:col-span-8 space-y-8">
                
                <div class="bg-white p-10 rounded-ucu border border-gray-100 shadow-sm flex flex-col md:flex-row items-center justify-between gap-6">
                    <div>
                        <h2 class="font-montserrat font-bold text-3xl text-ucu-foret leading-tight">Bonjour, <span class="text-ucu-lagon"><?= htmlspecialchars($particulier["nom"] ?? "Utilisateur") . " " . htmlspecialchars($particulier["prenom"] ?? "") ?></span> 👋</h2>
                        <p class="text-sm text-gray-400 mt-2 italic">Prêt(e) pour votre prochain geste écologique aujourd'hui ?</p>
                    </div>
                    <div class="flex gap-4">
                        <a href="particulier_annonces.php" class="bg-ucu-foret text-white px-6 py-3 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-lagon transition-all">Déposer</a>
                        <a href="particulier_catalogue.php" class="border-2 border-ucu-foret text-ucu-foret px-6 py-3 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-foret hover:text-white transition-all">Boutique</a>
                    </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div class="bg-ucu-foret p-8 rounded-ucu text-white group cursor-pointer hover:shadow-xl transition-all">
                        <span class="text-2xl mb-4 block">📦</span>
                        <h3 class="font-montserrat font-bold text-sm uppercase mb-2 tracking-widest text-ucu-lagon">Dépôt Conteneur</h3>
                        <p class="text-[10px] opacity-60 leading-relaxed mb-6">Demandez un code d'accès pour déposer vos objets en box sécurisée.</p>
                        <a href="particulier_conteneurs.php" class="inline-block text-[9px] font-bold uppercase border-b border-ucu-lagon pb-1">Suivre mes codes →</a>
                    </div>
                    
                    <div class="bg-white p-8 rounded-ucu border border-gray-100 group cursor-pointer hover:border-ucu-lagon transition-all">
                        <span class="text-2xl mb-4 block">📢</span>
                        <h3 class="font-montserrat font-bold text-sm uppercase mb-2 tracking-widest text-ucu-foret">Mes Annonces</h3>
                        <p class="text-[10px] text-gray-500 leading-relaxed mb-6 italic italic">Vendez ou donnez vos objets en main propre après validation.</p>
                        <a href="particulier_annonces.php" class="inline-block text-[9px] font-bold uppercase text-ucu-lagon border-b border-ucu-lagon pb-1">Gérer mes ventes →</a>
                    </div>
                </div>

                <div class="bg-white rounded-ucu border border-gray-100 overflow-hidden shadow-sm">
                    <div class="p-6 border-b border-gray-100 flex justify-between items-center">
                        <h4 class="font-montserrat font-bold text-[10px] uppercase tracking-widest text-ucu-foret">Dernières Activités</h4>
                        <span class="text-[9px] text-gray-400 font-bold uppercase italic">Mise à jour : Aujourd'hui</span>
                    </div>
                    <div class="p-6 space-y-4 text-xs">
                        <div class="flex items-center justify-between p-3 hover:bg-ucu-grey rounded-ucu transition-all">
                            <div class="flex items-center gap-4">
                                <span class="text-green-500">✔</span>
                                <p class="font-medium">Annonce "Chaises en bois" <span class="text-gray-400">validée</span></p>
                            </div>
                            <span class="text-[9px] text-gray-400 italic">Il y a 2h</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="lg:col-span-4 space-y-8">
                
                <div class="bg-white p-8 rounded-ucu border border-gray-100 shadow-sm text-center">
                    <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Votre Impact</h4>
                    <div class="relative w-32 h-32 mx-auto mb-6">
                        <svg class="w-full h-full transform -rotate-90">
                            <circle cx="64" cy="64" r="58" stroke="currentColor" stroke-width="8" class="text-ucu-grey" fill="transparent" />
                            <circle cx="64" cy="64" r="58" stroke="currentColor" stroke-width="8" class="text-ucu-lagon" fill="transparent" stroke-dasharray="364.4" stroke-dashoffset="91" stroke-linecap="round" />
                        </svg>
                        <div class="absolute inset-0 flex flex-col items-center justify-center">
                            <span class="text-3xl font-montserrat font-bold text-ucu-foret"><?= htmlspecialchars($particulier["score"] ?? 0) ?></span>
                            <span class="text-[8px] font-bold text-gray-400 uppercase tracking-widest">Points</span>
                        </div>
                    </div>
                    <p class="text-[10px] text-gray-400 italic leading-relaxed">Encore <?= 1000 - htmlspecialchars($particulier["score"] ?? 0) ?> points pour obtenir <br><span class="text-ucu-foret font-bold">10% sur votre prochain cours</span> !</p>
                </div>

                <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm">
                    <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Prochain RDV</h4>
                    <div class="flex gap-4 p-4 bg-ucu-grey rounded-ucu border-l-4 border-ucu-lagon">
                        <div class="text-center min-w-[30px]">
                            <p class="text-[9px] font-bold text-ucu-lagon">MAR</p>
                            <p class="text-xl font-montserrat font-bold text-ucu-foret">17</p>
                        </div>
                        <div>
                            <p class="text-[11px] font-bold text-ucu-foret uppercase">Soudure Artisanale</p>
                            <p class="text-[9px] text-gray-400">14:30 • Ivry</p>
                        </div>
                    </div>
                    <a href="particulier_planning.php" class="block text-center mt-6 text-[9px] font-bold text-ucu-lagon uppercase tracking-widest hover:underline">Voir mon calendrier complet</a>
                </div>

                <div class="bg-ucu-lagon p-6 rounded-ucu text-white shadow-md relative overflow-hidden">
                    <h4 class="font-montserrat font-bold text-xs uppercase mb-2">Espace Conseils</h4>
                    <p class="text-[10px] opacity-90 mb-4 leading-relaxed italic">Comment restaurer votre mobilier sans outils complexes ?</p>
                    <a href="particulier_conseils.php" class="bg-ucu-foret text-white px-4 py-2 rounded-ucu text-[9px] font-bold uppercase inline-block">Apprendre →</a>
                </div>

            </div>

        </div>
    </main>

</body>
</html>