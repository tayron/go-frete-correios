import http.client
import mimetypes
conn = http.client.HTTPSConnection("127.0.0.1", 3003)
dataList = []
boundary = 'wL36Yn8afVp8Ag7AmP8qZ0SA4n1v9T'
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=codigo_servido_desejado;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("41106")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=cep_origem;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("11680000")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=cep_destino;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("82220000")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=peso;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("1")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=altura;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("15")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=largura;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("22")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=comprimento;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("32")
dataList.append('--' + boundary)
dataList.append('Content-Disposition: form-data; name=valor_produto;')

dataList.append('Content-Type: {}'.format('multipart/form-data'))
dataList.append('')

dataList.append("0")
dataList.append('--'+boundary+'--')
dataList.append('')
body = '\r\n'.join(dataList)
payload = body
headers = {
   'Content-type': 'multipart/form-data; boundary={}'.format(boundary) 
}
conn.request("POST", "/frete", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))