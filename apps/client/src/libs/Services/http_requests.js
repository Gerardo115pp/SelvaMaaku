import { attributesToJson, HttpResponse } from "./base";
import { SVG_PREFIX, SALES_AICHAT_SERVICE } from "./addresses";

/*=============================================
=            Resources            =
=============================================*/

    export class GetSVGResourceRequest {
        constructor(filename) {
            this.filename = filename;
        }

        toJson = attributesToJson.bind(this);

        /**
         * Sends a request to the server and returns the response
         * @async
         * @returns {Promise<HttpResponse>}
         */
        do = async () => {
            const response = await fetch(`${SVG_PREFIX}/${this.filename}`);
            let data = null;

            if (response.status <= 200 && response.status < 300) {
                data = await response.text();
            }

            return new HttpResponse(response, data);
        }
    }

/*=====  End of Resources  ======*/


/*=============================================
=            Ai sales chat            =
=============================================*/

    
    /*----------  Authentication  ----------*/

        export class GetVerifyChatCredentialsRequest {
            /**
             * @typedef {Object} ChatCredentialsResponse
             * @property {boolean} response
             */

            /** @returns {Promise<HttpResponse<ChatCredentialsResponse>>} */
            do = async () => {
                const response = await fetch(`${SALES_AICHAT_SERVICE}/chat-claims/verify`);
                let data = null;

                if (response.ok) {
                    data = await response.json();
                }

                return new HttpResponse(response, data);
            }
        }

        export class GetChatCredentialsRequest {
            /**
             * @typedef {Object} ChatCredentialsResponse
             * @property {boolean} response
             */

            /** @returns {Promise<HttpResponse<ChatCredentialsResponse>>} */
            do = async () => {
                const response = await fetch(`${SALES_AICHAT_SERVICE}/chat-claims`);
                let data = null;

                if (response.ok) {
                    data = await response.json();
                }

                return new HttpResponse(response, data);
            }
        }

    /*----------  Chat room & messages  ----------*/

        export class GetChatRoomRequest {
            do = async () => {
                const response = await fetch(`${SALES_AICHAT_SERVICE}/chat`);

                let data = null;

                if (response.ok) {
                    data = await response.json();
                }

                return new HttpResponse(response, data);
            }
        }

        export class PostChatMessageRequest {
            constructor(content) {
                this.content = content;
            }

            toJson = attributesToJson.bind(this);

            do = async () => {
                const response = await fetch(`${SALES_AICHAT_SERVICE}/chat`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: this.toJson()
                });

                let data = null;

                if (response.ok) {
                    data = await response.json();
                }

                return new HttpResponse(response, data);
            }
        }
    

/*=====  End of Ai sales chat  ======*/

