const API_URL = "http://localhost:8080/terms";

let editandoId = null;

const campoTermo = document.querySelector("#termo");
const campoCategoria = document.querySelector("#categoria");
const campoExplicacao = document.querySelector("#explicacao");
const campoExemplo = document.querySelector("#exemplo");
const campoStatus = document.querySelector("#status");
const btnSalvar = document.querySelector("#btn-salvar");
const btnCancelar = document.querySelector("#btn-cancelar");

async function buscarTermos() {
  const resposta = await fetch(API_URL);
  const termos = await resposta.json();
  renderizarTermos(termos);
}

function criarCelula(texto) {
  const td = document.createElement("td");
  td.textContent = texto;
  return td;
}

function renderizarTermos(termos) {
  const tbody = document.querySelector("#tabela-termos tbody");
  tbody.innerHTML = "";

  for (const t of termos) {
    const tr = document.createElement("tr");
    tr.appendChild(criarCelula(t.termo));
    tr.appendChild(criarCelula(t.categoria));
    tr.appendChild(criarCelula(t.explicacao_simples));
    tr.appendChild(criarCelula(t.exemplo));
    tr.appendChild(criarCelula(t.status));

    const tdAcoes = document.createElement("td");

    const btnEditar = document.createElement("button");
    btnEditar.type = "button";
    btnEditar.textContent = "Editar";
    btnEditar.addEventListener("click", () => iniciarEdicao(t));
    tdAcoes.appendChild(btnEditar);

    const btnRemover = document.createElement("button");
    btnRemover.type = "button";
    btnRemover.textContent = "Remover";
    btnRemover.addEventListener("click", () => removerTermo(t.id));
    tdAcoes.appendChild(btnRemover);

    if (t.status === "estudando") {
      const btnEntendido = document.createElement("button");
      btnEntendido.type = "button";
      btnEntendido.textContent = "Marcar como entendido";
      btnEntendido.addEventListener("click", () => marcarComoEntendido(t));
      tdAcoes.appendChild(btnEntendido);
    }

    tr.appendChild(tdAcoes);
    tbody.appendChild(tr);
  }
}

function limparFormulario() {
  campoTermo.value = "";
  campoCategoria.value = "";
  campoExplicacao.value = "";
  campoExemplo.value = "";
  campoStatus.value = "estudando";
  editandoId = null;
  btnSalvar.textContent = "Salvar";
  btnCancelar.style.display = "none";
}

function iniciarEdicao(t) {
  campoTermo.value = t.termo;
  campoCategoria.value = t.categoria;
  campoExplicacao.value = t.explicacao_simples;
  campoExemplo.value = t.exemplo;
  campoStatus.value = t.status;
  editandoId = t.id;
  btnSalvar.textContent = "Atualizar";
  btnCancelar.style.display = "inline";
}

async function salvarTermo() {
  const termo = {
    termo: campoTermo.value,
    categoria: campoCategoria.value,
    explicacao_simples: campoExplicacao.value,
    exemplo: campoExemplo.value,
    status: campoStatus.value,
  };

  const url = editandoId === null ? API_URL : `${API_URL}/${editandoId}`;
  const metodo = editandoId === null ? "POST" : "PUT";

  const resposta = await fetch(url, {
    method: metodo,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(termo),
  });

  if (!resposta.ok) {
    const erro = await resposta.json();
    alert(erro.erro || "Erro ao salvar termo");
    return;
  }

  limparFormulario();
  buscarTermos();
}

async function removerTermo(id) {
  if (!confirm("Remover este termo?")) return;
  await fetch(`${API_URL}/${id}`, { method: "DELETE" });
  buscarTermos();
}

async function marcarComoEntendido(t) {
  const atualizado = { ...t, status: "entendido" };
  await fetch(`${API_URL}/${t.id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(atualizado),
  });
  buscarTermos();
}

btnSalvar.addEventListener("click", salvarTermo);
btnCancelar.addEventListener("click", limparFormulario);

buscarTermos();
