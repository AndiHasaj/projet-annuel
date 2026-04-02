<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Admin | Centre de Validation</title>
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

    <?php include '../includes/sidebar.php'; ?>

    <main class="lg:ml-72 p-12">
        
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter italic">Flux de <span class="text-ucu-lagon underline decoration-2 underline-offset-8">Validation</span></h2>
            <p class="text-sm text-gray-400 mt-4">Approuvez les propositions des salariés et les offres du catalogue avant publication.</p>
        </header>

        <div class="grid grid-cols-1 gap-8">
            
            <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                <div class="p-6 border-b border-gray-100 bg-ucu-grey/20">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Demandes d'activités (Salariés)</h3>
                </div>
                
                <div class="divide-y divide-gray-50">
                    <div class="p-8 flex flex-col md:flex-row justify-between items-center gap-6 hover:bg-ucu-grey/10 transition-all">
                        <div class="flex-1">
                            <div class="flex items-center gap-3 mb-2">
                                <span class="bg-ucu-lagon/10 text-ucu-lagon text-[8px] font-bold px-2 py-0.5 rounded uppercase">Formation</span>
                                <p class="text-[10px] text-gray-400 font-bold italic">Soumis par : LACI Senard</p>
                            </div>
                            <h4 class="font-montserrat font-bold text-sm text-ucu-foret uppercase">Atelier Soudure & Upcycling Industriel</h4>
                            <p class="text-xs text-gray-500 mt-2 leading-relaxed">Prévu le 25/03/2026 au Centre Ivry. Tarif : 45€ (Stripe).</p>
                        </div>
                        <div class="flex gap-3">
                            <button class="px-6 py-3 border border-gray-200 rounded-ucu text-[9px] font-bold uppercase text-gray-400 hover:bg-red-50 hover:text-red-500 transition-all">Refuser</button>
                            <button class="px-6 py-3 bg-ucu-foret text-white rounded-ucu text-[9px] font-bold uppercase tracking-widest hover:bg-ucu-lagon shadow-md transition-all">Valider l'offre</button>
                        </div>
                    </div>

                    <div class="p-8 flex flex-col md:flex-row justify-between items-center gap-6 hover:bg-ucu-grey/10 transition-all text-gray-400">
                        <div class="flex-1 opacity-60">
                            <div class="flex items-center gap-3 mb-2">
                                <span class="bg-gray-100 text-gray-400 text-[8px] font-bold px-2 py-0.5 rounded uppercase">Conseil technique</span>
                                <p class="text-[10px] font-bold italic">Soumis par : Andi HASAJ</p>
                            </div>
                            <h4 class="font-montserrat font-bold text-sm uppercase">Guide : Restauration de palettes PVC</h4>
                        </div>
                        <span class="text-[9px] font-bold uppercase italic italic tracking-widest">En cours de lecture...</span>
                    </div>
                </div>
            </section>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="bg-ucu-foret p-8 rounded-ucu text-white">
                    <h3 class="font-montserrat font-bold text-xs uppercase mb-6 tracking-widest text-ucu-lagon">Sécurité & Codes</h3>
                    <p class="text-[11px] opacity-70 mb-8 leading-relaxed italic">
                        Le système génère automatiquement des tokens d'accès uniques pour chaque artisan lors de la validation de leur contrat.
                    </p>
                    <button class="w-full bg-white text-ucu-foret py-4 rounded-ucu font-bold text-[9px] uppercase tracking-[0.2em] hover:bg-ucu-lagon hover:text-white transition-all">
                        Auditer les logs d'accès
                    </button>
                </div>

                <div class="bg-white p-8 rounded-ucu border border-gray-100 shadow-sm">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Documents en attente</h3>
                    <div class="space-y-4">
                        <div class="flex justify-between items-center p-3 bg-ucu-grey rounded">
                            <span class="text-[10px] font-bold text-ucu-foret">Contrat_Artisan_78.pdf</span>
                            <button class="text-[18px]">📥</button>
                        </div>
                        <div class="flex justify-between items-center p-3 bg-ucu-grey rounded">
                            <span class="text-[10px] font-bold text-ucu-foret">Bilan_Financier_T1.pdf</span>
                            <button class="text-[18px]">📥</button>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Admin • Système Informatique Centralisé • 2026</p>
    </footer>

</body>
</html>