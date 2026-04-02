<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Pro | Gestion des Projets</title>
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

    <nav class="bg-ucu-foret border-b border-white/10 sticky top-0 z-50 text-white">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a href="pro_dashboard.php" class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-foret font-montserrat font-bold italic">PRO</a>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest">Portfolio Projets</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-white/60">
                <a href="pro_dashboard.php" class="hover:text-white">Dashboard</a>
                <a href="pro_projets.php" class="text-ucu-lagon">Mes Projets</a>
                <a href="pro_abonnements.php" class="hover:text-white">Abonnement</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <header class="flex justify-between items-end mb-12">
            <div>
                <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Suivi des <span class="text-ucu-lagon italic">Créations</span></h2>
                <p class="text-sm text-gray-400 mt-2 text-balance">Documentez vos transformations pour booster votre visibilité auprès des particuliers.</p>
            </div>
            <button class="bg-ucu-foret text-white px-8 py-4 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-lagon shadow-lg transition-all">
                + Nouveau Projet
            </button>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            
            <div class="bg-white rounded-ucu border border-gray-100 overflow-hidden shadow-sm group">
                <div class="h-48 bg-ucu-grey relative">
                    <div class="absolute top-4 left-4 bg-orange-400 text-white text-[8px] font-bold px-3 py-1 rounded-full uppercase">En cours : Finition</div>
                    <div class="w-full h-full flex items-center justify-center text-5xl">🔨</div>
                </div>
                <div class="p-6">
                    <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mb-2">Buffet Scandinave "Néo"</h3>
                    <p class="text-[10px] text-gray-400 italic mb-6">Matière : Chêne récupéré (Annonce #492)</p>
                    
                    <div class="space-y-4">
                        <div class="w-full bg-gray-100 h-1.5 rounded-full overflow-hidden">
                            <div class="bg-ucu-lagon h-full w-[75%]"></div>
                        </div>
                        <div class="flex justify-between items-center pt-4 border-t border-gray-50">
                            <span class="text-[9px] font-bold text-gray-400 uppercase">Étape 3/4</span>
                            <button class="text-ucu-foret font-bold text-[9px] uppercase hover:underline">Mettre à jour</button>
                        </div>
                    </div>
                </div>
            </div>

            <div class="bg-white rounded-ucu border-2 border-ucu-lagon overflow-hidden shadow-md relative">
                <div class="absolute top-4 right-4 z-10">
                    <span class="bg-ucu-lagon text-white text-[8px] font-bold px-3 py-1 rounded-full uppercase shadow-sm">🚀 Mis en avant</span>
                </div>
                <div class="h-48 bg-ucu-grey flex items-center justify-center text-5xl">✨</div>
                <div class="p-6">
                    <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mb-2 text-balance">Lampe Industrielle "Volt"</h3>
                    <p class="text-[10px] text-gray-400 italic mb-6">Matière : Cuivre & Métal (Box Ivry)</p>
                    
                    <div class="flex items-center justify-between pt-4 border-t border-gray-50">
                        <div class="flex gap-2">
                            <span class="text-[14px]">👁️ 142</span>
                            <span class="text-[14px]">❤️ 28</span>
                        </div>
                        <button class="bg-ucu-foret text-white px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-ucu-lagon transition-all">Voir sur le site</button>
                    </div>
                </div>
            </div>

            <div class="bg-white rounded-ucu border border-gray-100 overflow-hidden shadow-sm opacity-60">
                <div class="h-48 bg-gray-100 flex items-center justify-center text-5xl grayscale">📦</div>
                <div class="p-6">
                    <h3 class="font-montserrat font-bold text-gray-400 text-sm uppercase mb-2">Table basse Résine</h3>
                    <p class="text-[10px] text-gray-300 italic mb-6">Vendu le 02/03/2026</p>
                    <div class="pt-4 border-t border-gray-50">
                        <span class="text-[9px] font-bold text-gray-300 uppercase italic">Vendu (Score +500)</span>
                    </div>
                </div>
            </div>

        </div>

        <section class="mt-16 bg-ucu-foret p-10 rounded-ucu text-white flex flex-col md:flex-row items-center gap-8">
            <div class="flex-1">
                <h4 class="font-montserrat font-bold text-lg uppercase mb-2 tracking-tighter">Besoin de plus de visibilité ?</h4>
                <p class="text-xs opacity-70 italic leading-relaxed">Activez une campagne publicitaire pour que vos projets apparaissent en priorité dans le catalogue des particuliers.</p>
            </div>
            <a href="pro_abonnements.php" class="bg-white text-ucu-foret px-8 py-4 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-lagon hover:text-white transition-all">Booster un projet</a>
        </section>

    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Pro • Portfolio Management • 2026</p>
    </footer>

</body>
</html>