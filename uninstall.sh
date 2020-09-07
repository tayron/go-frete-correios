#!/bin/bash
diretorioAplicacao="/usr/local/gocep/"

nomeServico="go-frete-correios"
nomeArquivoServico="go-frete-correios.service"
caminhoScriptServicoLinux="/etc/systemd/system/"

echo "Desinstalando GO-FRETE - Serviço de consulta de endereço por FRETE"
echo 

if service $nomeServico stop; then
    echo "[√] Parando serviço"
else
    echo
    echo "[!] Falha ao para o serviço"
    echo

    exit 1    
fi

if [ -e $diretorioAplicacao ] && rm -R $diretorioAplicacao; then
    echo "[√] Removendo instalação antiga"
fi

if [ -e "$caminhoScriptServicoLinux$nomeArquivoServico" ] && rm "$caminhoScriptServicoLinux$nomeArquivoServico"; then
    echo "[√] Removendo serviço antigo"
fi

if systemctl daemon-reload; then 
    echo "[√] Atualizando lista de serviços do sistema operacional"
else 
    echo
    echo "[!] Erro ao atualizar lista de serviços do sistema operacional"
    echo

    exit 1    
fi

echo 
echo "Desinstalação concluída"