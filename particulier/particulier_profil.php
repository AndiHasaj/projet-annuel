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
    <title>UCU | Mon Profil & Score</title>
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
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">

    <nav class="bg-white border-b border-gray-100 sticky top-0 z-50">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a href="particulier_dashboard.php" class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold italic">UCU</a>
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Mon Profil</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-gray-400 hover:text-ucu-foret">Catalogue</a>
                <a href="particulier_conseils.php" class="text-gray-400 hover:text-ucu-foret">Conseils</a>
                <a href="particulier_annonces.php" class="text-gray-400 hover:text-ucu-foret">Annonces</a>
                    <a href="particulier_conteneurs.php" class="text-gray-400 hover:text-ucu-foret">Conteneurs</a>
                    <a href="particulier_planning.php" class="text-gray-400 hover:text-ucu-foret">Planning</a>
                <a href="particulier_profil.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-12">
            
            <div class="space-y-8">
                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100 text-center">
                    <div class="w-24 h-24 bg-ucu-grey rounded-full mx-auto mb-4 flex items-center justify-center text-3xl border-2 border-ucu-lagon">👤</div>
                    <h2 class="font-montserrat font-bold text-ucu-foret uppercase text-lg italic"><?= htmlspecialchars($particulier["nom"] ?? "Utilisateur") . " " . htmlspecialchars($particulier["prenom"] ?? "") ?></h2>
                    <p class="text-[10px] text-gray-400 font-bold uppercase tracking-[0.2em] mb-6">Membre depuis <?= htmlspecialchars($particulier["created_at"] ?? "Mars 2026") ?></p>
                    
                    <div class="space-y-3 text-left">
                        <div class="p-3 bg-ucu-grey rounded-ucu">
                            <p class="text-[8px] font-bold text-gray-400 uppercase">Email</p>
                            <p class="text-xs font-medium text-ucu-foret"><?= htmlspecialchars($particulier["email"] ?? "") ?></p>
                        </div>
                        <div class="p-3 bg-ucu-grey rounded-ucu">
                            <p class="text-[8px] font-bold text-gray-400 uppercase">Ville</p>
                            <p class="text-xs font-medium text-ucu-foret"><?= htmlspecialchars($particulier["ville"] ?? "Ivry-sur-Seine") ?></p>
                        </div>
                    </div>
                    <button class="w-full mt-6 border-2 border-ucu-foret text-ucu-foret py-3 rounded-ucu font-bold text-[9px] uppercase tracking-widest hover:bg-ucu-foret hover:text-white transition-all">Modifier mes infos</button>
                </section>

                <section class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest italic">Mes Badges</h3>
                    <div class="flex flex-wrap gap-4 justify-center">
                        <div class="w-12 h-12 bg-ucu-lagon/20 rounded-full flex items-center justify-center text-xl grayscale-0" title="Premier Don">🎁</div>
                        <div class="w-12 h-12 bg-ucu-grey rounded-full flex items-center justify-center text-xl grayscale" title="Expert Bois">🪚</div>
                        <div class="w-12 h-12 bg-ucu-grey rounded-full flex items-center justify-center text-xl grayscale" title="Éco-Acheteur">💳</div>
                    </div>
                </section>
            </div>

            <div class="lg:col-span-2 space-y-8">
                
                <section class="bg-ucu-foret p-10 rounded-ucu shadow-xl text-white">
                    <div class="flex flex-col md:flex-row justify-between items-center gap-8">
                        <div class="text-center md:text-left">
                            <h3 class="font-montserrat font-bold text-2xl uppercase tracking-tighter mb-2">Upcycling <span class="text-ucu-lagon">Score</span></h3>
                            <p class="text-xs opacity-70 italic mb-6 leading-relaxed">Votre score reflète votre engagement pour la planète. Continuez ainsi !</p>
                            <div class="flex gap-4">
                                <div class="bg-white/10 p-4 rounded-ucu">
                                    <p class="text-2xl font-montserrat font-bold text-ucu-lagon"><?= $particulier["score"] ?? 0 ?></p>
                                    <p class="text-[8px] uppercase font-bold opacity-60">Points cumulés</p>
                                </div>
                                <div class="bg-white/10 p-4 rounded-ucu">
                                    <p class="text-2xl font-montserrat font-bold text-ucu-lagon">14kg</p>
                                    <p class="text-[8px] uppercase font-bold opacity-60">Matière sauvée</p>
                                </div>
                            </div>
                        </div>
                        <div class="relative w-40 h-40">
                             <svg class="w-full h-full transform -rotate-90">
                                <circle cx="80" cy="80" r="70" stroke="currentColor" stroke-width="10" class="text-white/10" fill="transparent" />
                                <circle cx="80" cy="80" r="70" stroke="currentColor" stroke-width="10" class="text-ucu-lagon" fill="transparent" stroke-dasharray="440" stroke-dashoffset="110" stroke-linecap="round" />
                            </svg>
                            <div class="absolute inset-0 flex flex-col items-center justify-center">
                                <span class="text-xs font-bold uppercase tracking-widest text-ucu-lagon italic">Niveau</span>
                                <span class="text-2xl font-montserrat font-bold uppercase">Initié</span>
                            </div>
                        </div>
                    </div>
                </section>

                <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-6 border-b border-gray-100">
                        <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest italic">Historique & Documents</h4>
                    </div>
                    <table class="w-full text-left text-[11px]">
                        <thead class="bg-ucu-grey text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                            <tr>
                                <th class="px-6 py-4 italic">Action</th>
                                <th class="px-6 py-4">Date</th>
                                <th class="px-6 py-4">Score</th>
                                <th class="px-6 py-4 text-right">Document</th>
                            </tr>
                        </thead>
                        <tbody class="divide-y divide-gray-50">
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5">
                                    <p class="font-bold text-ucu-foret uppercase">Formation Soudure</p>
                                    <p class="text-[9px] text-gray-400 italic">Paiement Stripe confirmé</p>
                                </td>
                                <td class="px-6 py-5 text-gray-500">14/03/2026</td>
                                <td class="px-6 py-5 text-green-500 font-bold">+100 pts</td>
                                <td class="px-6 py-5 text-right">
                                    <button class="bg-ucu-foret text-white px-3 py-1.5 rounded-ucu font-bold uppercase text-[8px] hover:bg-ucu-lagon transition-all">📄 Facture PDF</button>
                                </td>
                            </tr>
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5">
                                    <p class="font-bold text-ucu-foret uppercase">Don : Chaises Vintage</p>
                                    <p class="text-[9px] text-gray-400 italic">Récupéré par un artisan</p>
                                </td>
                                <td class="px-6 py-5 text-gray-500">08/03/2026</td>
                                <td class="px-6 py-5 text-green-500 font-bold">+250 pts</td>
                                <td class="px-6 py-5 text-right">
                                    <button class="bg-ucu-foret text-white px-3 py-1.5 rounded-ucu font-bold uppercase text-[8px] hover:bg-ucu-lagon transition-all">📜 Attestation</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </section>

                <section class="bg-white p-8 rounded-ucu border-2 border-dashed border-ucu-grey">
                    <div class="flex items-center justify-between">
                        <div>
                            <h4 class="font-montserrat font-bold text-xs uppercase text-ucu-foret mb-1">Notifications Push</h4>
                            <p class="text-[10px] text-gray-400 italic italic">Géré via OneSignal pour ne manquer aucune validation.</p>
                        </div>
                        <label class="relative inline-flex items-center cursor-pointer">
                            <input type="checkbox" value="" class="sr-only peer" checked>
                            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-ucu-lagon"></div>
                        </label>
                    </div>
                </section>

            </div>
        </div>
    </main>

</body>
</html>