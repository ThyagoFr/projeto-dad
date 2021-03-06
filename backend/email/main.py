from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import smtplib
 

def lambda_handler(event, context):

    to = event["Records"][0]["messageAttributes"]["To"]["stringValue"]
    name = event["Records"][0]["messageAttributes"]["Name"]["stringValue"]
    token = event["Records"][0]["messageAttributes"]["Token"]["stringValue"]
    
    html =  """\
            <html>
            <head></head>
            <body>
                <p>Olá, {0}<br>
                Você solicitou a recuperação de senha na plataforma BooksDAD . <br>
                Sua nova senha temporária é <b> {1} </b>. Ao entrar na plataforma, por favor atualizar sua senha para uma de sua escolha.
                </p>
            </body>
            </html>
            """.format(name, token)

    msg = MIMEMultipart("alternative")
    password = "392035ab"
    msg['From'] = "bookdad2020@gmail.com"
    msg['To'] = to
    msg['Subject'] = "Recuperação de senha"
    msg.attach(MIMEText(html, 'html'))
    server = smtplib.SMTP('smtp.gmail.com: 587')
    server.starttls()
    server.login(msg['From'], password)
    server.sendmail(msg['From'], msg['To'], msg.as_string())
    server.quit()