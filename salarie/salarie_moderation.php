<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Salarié | Modération Forum</title>
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

    <nav class="bg-ucu-lagon border-b border-white/10 sticky top-0 z-50 text-white">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-lagon font-montserrat font-bold italic">SAL</div>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest text-ucu-foret">Espace Salarié</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-ucu-foret">
                <a href="salarie_dashboard.php" class="hover:text-white transition-colors">Dashboard</a>
                <a href="salarie_evenements.php" class="hover:text-white transition-colors">Événements</a>
                <a href="salarie_moderation.php" class="border-b-2 border-ucu-foret pb-1">Forums</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter italic">Modération <span class="text-white bg-ucu-lagon px-2 rounded">Forums</span></h2>
            <p class="text-sm text-gray-400 mt-2">Gérez les échanges et garantissez la qualité des conseils communautaires.</p>
        </header>

        <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
            
            <aside class="lg:col-span-1 space-y-4">
                <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100">
                    <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Flux à traiter</h4>
                    <nav class="space-y-2">
                        <button class="w-full text-left px-4 py-3 bg-ucu-grey rounded-ucu text-[11px] font-bold text-red-500 flex justify-between">
                            🚩 Signalements <span class="bg-red-100 px-2 py-0.5 rounded">3</span>
                        </button>
                        <button class="w-full text-left px-4 py-3 hover:bg-ucu-grey rounded-ucu text-[11px] font-bold text-gray-400 flex justify-between">
                            💬 Questions <span class="bg-gray-100 px-2 py-0.5 rounded">12</span>
                        </button>
                        <button class="w-full text-left px-4 py-3 hover:bg-ucu-grey rounded-ucu text-[11px] font-bold text-gray-400">
                            🛡️ Messages archivés
                        </button>
                    </nav>
                </div>
            </aside>

            <div class="lg:col-span-3 space-y-6">
                
                <div class="bg-white rounded-ucu shadow-sm border-l-4 border-red-500 overflow-hidden">
                    <div class="p-6">
                        <div class="flex justify-between items-start mb-4">
                            <div class="flex gap-3 items-center">
                                <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center text-xs">👤</div>
                                <div>
                                    <p class="text-[11px] font-bold text-ucu-foret uppercase tracking-tighter">Jean_D (Particulier)</p>
                                    <p class="text-[9px] text-gray-400 italic">Posté sur : Tuto Soudure Arc</p>
                                </div>
                            </div>
                            <span class="text-[8px] bg-red-50 text-red-500 font-bold px-2 py-1 rounded uppercase">Signalé par 2 membres</span>
                        </div>
                        <p class="text-xs text-gray-600 bg-ucu-grey p-4 rounded-ucu italic mb-6">
                            "Si vous cherchez du matériel pas cher, allez voir sur le site [Lien_Spam].com, c'est bien mieux que ce que propose l'artisan !"
                        </p>
                        <div class="flex gap-3">
                            <button class="bg-ucu-foret text-white px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-red-600 transition-all">Supprimer</button>
                            <button class="border border-gray-200 text-gray-400 px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-gray-50 transition-all">Ignorer</button>
                            <button class="text-[9px] font-bold text-ucu-lagon uppercase ml-auto hover:underline">Voir l'historique du membre</button>
                        </div>
                    </div>
                </div>

                <div class="bg-white rounded-ucu shadow-sm border-l-4 border-ucu-lagon overflow-hidden">
                    <div class="p-6">
                        <div class="flex justify-between items-start mb-4">
                            <div class="flex gap-3 items-center">
                                <div class="w-8 h-8 bg-ucu-lagon/10 rounded-full flex items-center justify-center text-xs text-ucu-lagon">❓</div>
                                <div>
                                    <p class="text-[11px] font-bold text-ucu-foret uppercase tracking-tighter">Alice_V (Particulier)</p>
                                    <p class="text-[9px] text-gray-400 italic">Question technique • 1h ago</p>
                                </div>
                            </div>
                            <span class="text-[8px] bg-ucu-lagon/10 text-ucu-lagon font-bold px-2 py-1 rounded uppercase tracking-widest">Urgent</span>
                        </div>
                        <p class="text-xs text-gray-600 mb-6 font-medium">
                            "Bonjour, pour le tutoriel sur la rénovation de chaises, peut-on utiliser de l'huile de lin directement ou faut-il la diluer avec de l'essence de térébenthine ?"
                        </p>
                        <div class="flex gap-3">
                            <textarea placeholder="Votre réponse officielle en tant que formateur..." class="flex-1 bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:ring-1 focus:ring-ucu-lagon"></textarea>
                        </div>
                        <div class="mt-4 flex justify-end">
                            <button class="bg-ucu-lagon text-white px-6 py-2 rounded-ucu font-bold text-[9px] uppercase tracking-widest shadow-md">Répondre</button>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Salarié • Console de Modération • 2026</p>
    </footer>

</body>
</html>