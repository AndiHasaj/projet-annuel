<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Pro | Logistique & Collecte</title>
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
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest">Logistique & Collecte</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-white/60">
                <a href="pro_dashboard.php" class="hover:text-white">Dashboard</a>
                <a href="pro_marche.php" class="hover:text-white">Marché</a>
                <a href="pro_logistique.php" class="text-ucu-lagon">Logistique</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            
            <section class="lg:col-span-5 space-y-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h2 class="font-montserrat font-bold text-lg text-ucu-foret uppercase mb-6 italic tracking-tight">Valider une récupération</h2>
                    
                    <div class="space-y-6">
                        <div class="aspect-square bg-ucu-grey rounded-ucu border-2 border-dashed border-gray-200 flex flex-col items-center justify-center relative overflow-hidden group cursor-pointer">
                            <div class="absolute inset-0 bg-gradient-to-t from-ucu-foret/20 to-transparent"></div>
                            <span class="text-4xl mb-4 z-10">📷</span>
                            <p class="text-[10px] font-bold uppercase text-ucu-foret z-10">Scanner le QR Code</p>
                            <p class="text-[8px] text-gray-400 mt-1 z-10 italic">Placez le code de l'objet dans le cadre</p>
                            <div class="absolute top-0 left-0 w-full h-1 bg-ucu-lagon shadow-[0_0_15px_#6A9A9A] animate-pulse"></div>
                        </div>

                        <div class="relative">
                            <div class="absolute inset-0 flex items-center"><span class="w-full border-t border-gray-100"></span></div>
                            <div class="relative flex justify-center text-[10px] uppercase"><span class="bg-white px-2 text-gray-300 font-bold tracking-widest">OU SAISIE MANUELLE</span></div>
                        </div>

                        <form class="space-y-4">
                            <div>
                                <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block tracking-widest">Référence de l'objet</label>
                                <input type="text" placeholder="Ex: UCU-8492-BX" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none border border-transparent focus:border-ucu-lagon font-mono tracking-widest">
                            </div>
                            <button type="submit" class="w-full bg-ucu-foret text-white py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-[0.2em] shadow-lg hover:bg-ucu-lagon transition-all">
                                Confirmer la collecte
                            </button>
                        </form>
                    </div>
                </div>

                <div class="bg-ucu-lagon p-6 rounded-ucu text-white">
                    <p class="text-[10px] font-bold uppercase mb-2">💡 Note Pro</p>
                    <p class="text-[10px] leading-relaxed opacity-90 italic">
                        La validation de collecte génère automatiquement une notification Push au donateur et crédite ses points. Assurez-vous que l'objet est conforme à la photo de l'annonce avant de valider.
                    </p>
                </div>
            </section>

            <section class="lg:col-span-7">
                <div class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-6 border-b border-gray-100 flex justify-between items-center bg-ucu-grey/30">
                        <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Flux Logistique Récent</h3>
                        <span class="text-[9px] bg-white border border-gray-200 px-3 py-1 rounded-full text-gray-400 font-bold italic">Ivry-sur-Seine</span>
                    </div>
                    
                    <div class="divide-y divide-gray-50">
                        <div class="p-6 flex items-center justify-between hover:bg-ucu-grey/20 transition-all">
                            <div class="flex items-center gap-4">
                                <div class="w-12 h-12 bg-green-50 rounded-ucu flex items-center justify-center text-xl">✅</div>
                                <div>
                                    <p class="text-[11px] font-bold text-ucu-foret uppercase">Armoire de bureau métal</p>
                                    <p class="text-[9px] text-gray-400 italic">Récupéré aujourd'hui à 14:32 • Box #04</p>
                                </div>
                            </div>
                            <div class="text-right">
                                <p class="text-[10px] font-bold text-ucu-lagon">+12kg</p>
                                <p class="text-[8px] text-gray-300 uppercase font-bold tracking-tighter">Stock ajouté</p>
                            </div>
                        </div>

                        <div class="p-6 flex items-center justify-between hover:bg-ucu-grey/20 transition-all">
                            <div class="flex items-center gap-4">
                                <div class="w-12 h-12 bg-blue-50 rounded-ucu flex items-center justify-center text-xl">🚚</div>
                                <div>
                                    <p class="text-[11px] font-bold text-ucu-foret uppercase">Lot de planches chêne</p>
                                    <p class="text-[9px] text-gray-400 italic">Prévu demain • Box #12</p>
                                </div>
                            </div>
                            <button class="text-[8px] font-bold text-ucu-foret uppercase border border-ucu-foret px-3 py-1.5 rounded-ucu hover:bg-ucu-foret hover:text-white transition-all">
                                Voir Itinéraire
                            </button>
                        </div>

                        <div class="p-6 flex items-center justify-between opacity-50 grayscale">
                            <div class="flex items-center gap-4">
                                <div class="w-12 h-12 bg-gray-100 rounded-ucu flex items-center justify-center text-xl">📦</div>
                                <div>
                                    <p class="text-[11px] font-bold text-ucu-foret uppercase">Chutes textile (Coton)</p>
                                    <p class="text-[9px] text-gray-400 italic italic">Collecté le 12/03/2026</p>
                                </div>
                            </div>
                            <span class="text-[9px] font-bold uppercase text-gray-400 italic">Archivé</span>
                        </div>
                    </div>
                </div>
            </section>

        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Pro • Flux Logistique • v1.0</p>
    </footer>

</body>
</html>