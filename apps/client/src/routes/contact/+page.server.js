import nodemailer from 'nodemailer';
import Handlebars from 'handlebars';
import fs from 'fs'
import { info } from 'console';

/**
 * @type {import('./$types').Actions}
 */
export const actions = {
    submit: async http_request => {
        const form_data = await http_request.request.formData();

        const name = form_data.get('name');
        const email = form_data.get('email');
        const phone = form_data.get('phone');
        const message = form_data.get('message');

        console.log(form_data);

        if (process.env.MAIL_USERNAME === undefined || process.env.MAIL_PASSWORD === undefined || process.env.MAIL_PORT === undefined) {
            console.warn("someone submitted a contact form but the mail server is not configured.")
            return {
                status: 500,
                body: {
                    error: "Mail server not configured."
                }
            }
        }       

        const transporter = nodemailer.createTransport({
            service: 'gmail',
            auth: {
                user: process.env.MAIL_USERNAME,
                pass: process.env.MAIL_PASSWORD
            },
            port: parseInt(process.env.MAIL_PORT),
        });

        const contact_template_path = `${process.env.NODE_SERVER_OD}/email/contact_form_template.html`
        let html_body = "";

        if (fs.existsSync(contact_template_path)) {
            const html_source = fs.readFileSync(contact_template_path, "utf8");
            const template = Handlebars.compile(html_source);
            
            html_body = template({ name, email, phone, message });
        } else {
            console.warn(`Contact form template not found at: ${contact_template_path}`)
        }


        const mail_options = {
            from: `Selva Maaku Contact <${process.env.MAIL_USERNAME}>`,
            to: process.env.MAIL_CONTACT_EMAIL,
            subject: `Contact form submission from ${name}`,
            text: `Name: ${name}\nEmail: ${email}\nPhone: ${phone}\nMessage: ${message}`,
            html: html_body,
        }

        try {
            const resolution = transporter.sendMail(mail_options);

            resolution.then(info => {
                console.log(info);
            });
            

            return {
                status: 200,
                success: true,
            }
        } catch (error) {
            console.error(error);

            return {
                status: 500,
                success: false,
            }
        }
    }
}