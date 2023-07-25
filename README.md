# DDD - Domain Driven Design

### O que é DDD?
É uma forma de desenvolver software com o foco no coração da aplicação, o que chamamos de domínio,
tendo o objetivo de entender suas regras, processos e complexidades,
separando-as assim de outros pontos complexos que normalmente são adicionados durante o processo de desenvolvimento.


### Complexidade de Software
- DDD é e deve ser aplicado para casos de projetos de softwares complexos
- Grandes projetos possuem muitas áreas, muitas regras de negócio, muitas pessoas com diferentes visões em diferentes contextos.
- Não há como não utilizar técnicas avançadas em projetos de alta complexidade.
- Grande parte da complexidade desse tipo de software não vem da tecnologia, mas sim da comunicação, separação de contextos, entendimento do negócio por diversos ângulos.
- Pessoas


### Como o DDD pode nos ajudar?
 - Entender com profundidade o domínio e subdomínios da aplicação
 - Ter uma linguagem universal entre todos os envolvidos - (Linguagem Ubíqua)
 - Criar o design estratégico utilizando o Bounded Contexts
 - Criar o design tático para conseguir mapear e agregar as entidades e objetos de valor da aplicação, bem como os eventos do domínio.
 - Clareza do que é complexidade de negócio e complexidade técnica.


#### algumas técnicas que podem nos ajudar
Normalmente, cada setor de uma empresa tem sua própria "linguagem", essa linguagem nos ajuda a identificar os contextos, quando essa linguagem
começa a mudar podemos perceber que ali existe um limite entre um setor e outro, podemos então começar a delimitar os contextos.

# Domínio e Subdomínios
Dominio
- A grosso modo, podemos dizer que domínio é o ramo da empresa. obviamente que o software irá "resolver".


Core Domain
- Dominio principal
- Diferencial competitivo da empresa

Suport subdomain
- Apoiam o domínio
- Ajuda a atingir o objetibo do domain.

Generic subdomain
- Softwares auxiliares
- Sem diferencial competitivo
- Facilmente substituído

# Problema vs Solução

### Problema
- Visão geral do domínio e suas complexidades
- Onde fazemos a separação do problema em subdomínios
- Enxergamos, compreendemos o problema

### Solução
- Análise e modelagem do domínio
- Onde delimitamos os contextos
- Resolvemos o problema

## Contexto é Rei
se tem duas palavras iguais, mas que representam coisas direrentes, então temos contextos diferente. Se temos contexto
diferentes, teremos que criar entidades diferentes, mesmo que essas entidades sejam identicas.
Porque ainda que sejam identicas, elas respondem a contextos diferentes e podem mudar por coisas diferentes e o efeito colateral
ao alterar a entidade, se a entidade for unica o efeito colateral vai ser propagado para contextos diferentes...


Se temos duas palavras diferentes, mas que significam a mesma coisa, então provavelmente temos contextos diferentes.

# Visão Estratégica

## Context Mapping
 - Partnership
 - Shared Kernel
 - Customer-Supplier Development
 - Conformist
 - Anticorruption-layer
 - Open host service
 - Published ways
 - Big Ball of Mud

### Links
- CONTEXT MAPPRING: https://github.com/ddd-crew/context-mapping
- DDD: https://github.com/ddd-crew
