-- Criação das tabelas principais
CREATE TABLE IF NOT EXISTS shopping_lists (
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS list_items (
    id VARCHAR(50) PRIMARY KEY,
    shopping_list_id VARCHAR(50) NOT NULL REFERENCES shopping_lists(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    category_id VARCHAR(50),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS tips (
    id VARCHAR(50) PRIMARY KEY,
    text TEXT NOT NULL
);

-- Dados de exemplo para categorias
INSERT INTO categories (id, name) VALUES
  ('1', 'Alimentos'),
  ('2', 'Limpeza')
ON CONFLICT (id) DO NOTHING;

-- Dados de exemplo para dicas
INSERT INTO tips (id, text) VALUES
  ('1', 'Compare preços antes de comprar.'),
  ('2', 'Prefira produtos da estação.')
ON CONFLICT (id) DO NOTHING; 