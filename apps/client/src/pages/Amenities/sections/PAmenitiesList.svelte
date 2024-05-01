<script>
    import { supported_languages, site_language } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * @typedef {Object} SelvaAmenity
         * @property {string} title_en
         * @property {string} title_es
         * @property {string} description_en
         * @property {string} description_es
         */
        

        /** @type {SelvaAmenity[]} */
        const selva_maaku_amenities = [
            {
                title_en: "Security",
                title_es: "Seguridad",
                description_en: "Controlled access 24/7 – 365 with guardhouse",
                description_es: "Acceso controlado 24/7 – 365 con caseta de vigilancia"
            },
            {
                title_en: "Infinity Pool",
                title_es: "Alberca infinita",
                description_en: "Club house with pool and multipurpose palapa",
                description_es: "Casa club con alberca y palapa multipropósito"
            },
            {
                title_en: "Pet Friendly",
                title_es: "Pet Friendly",
                description_en: "100% pet friendly development",
                description_es: "Desarrollo 100% pet friendly"
            },
            {
                title_en: "Garbage separation",
                title_es: "Separación de basura",
                description_en: "Embrace sustainability with our dedicated garbage separation system, designed for eco-conscious living in our community",
                description_es: "Abraza la sostenibilidad con nuestro sistema de separación de basura dedicado, diseñado para una vida eco-consciente en nuestra comunidad"
            },
            {
                title_en: "Resident care management",
                title_es: "Administración de cuidado de residentes",
                description_en: "Experience unparalleled support ensuring a harmonious and well-maintained living environment for all",
                description_es: "Experimenta un apoyo sin igual que garantiza un entorno de vida armonioso y bien mantenido para todos"
            }
        ];
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="smk-ap-amenities-list">
    <ul id="smk-apal-content-wrapper" class="design-content-width">
        {#each selva_maaku_amenities as amenity}
            <li class="smk-amenity-item">
                <h3 class="smk-ai-title">
                    {#if $site_language !== supported_languages.SPANISH}
                        {amenity.title_en}
                    {:else}
                        {amenity.title_es}
                    {/if}
                </h3>
                <p class="smk-ai-description">
                    {#if $site_language !== supported_languages.SPANISH}
                        {amenity.description_en}
                    {:else}
                        {amenity.description_es}
                    {/if}
                </p>
            </li>
        {/each}
    </ul>
</section>

<style>
    #smk-ap-amenities-list {
        display: flex;
        flex-direction: column;
        align-items: center;
        container-type: inline-size;
        padding: var(--spacing-5) 0;
    }

    ul#smk-apal-content-wrapper {
        display: flex;
        padding: 0 var(--spacing-2);
        justify-content: center;
        flex-wrap: wrap;
        row-gap: var(--spacing-4);
        column-gap: var(--spacing-2);
    }

    li.smk-amenity-item {
        display: flex;
        width: min(384px, 31.57894%);
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-2);
        text-align: center;

        & h3.smk-ai-title {
            font-size: var(--font-size-h5);
        }   
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-ap-amenities-list {
                padding: var(--spacing-5) var(--spacing-2);
            }

            ul#smk-apal-content-wrapper {
                flex-direction: column;
                row-gap: var(--spacing-4);
                column-gap: 0;
            }

            li.smk-amenity-item {
                width: 100%;
            }

            li.smk-amenity-item h3.smk-ai-title {
                font-size: calc(var(--font-size-h2) * 1.2);
                white-space: nowrap;
            }
        }

        @media only screen and (max-width: 480px) {
            li.smk-amenity-item h3.smk-ai-title {
                white-space: pre-wrap !important;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>
