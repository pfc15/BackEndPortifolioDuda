CREATE TABLE IF NOT EXISTS Foto(
    id INTEGER PRIMARY KEY,
    Titulo TEXT UNIQUE,
    path_foto TEXT UNIQUE,
    Descricao TEXT
);
CREATE INDEX IF NOT EXISTS idx_titulo_foto
    ON Foto(Titulo);

CREATE TABLE IF NOT EXISTS tema (
    id INTEGER PRIMARY KEY,
    titulo TEXT UNIQUE,
    Foto INTERGER UNIQUE,
    ordem INTEGER UNIQUE,
    FOREIGN KEY (Foto) REFERENCES Foto(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS  Obra (
    id INTEGER PRIMARY KEY,
    titulo TEXT UNIQUE,
    Foto INTEGER UNIQUE,
    periodo DATE,
    descricao TEXT,
    ordem INTEGER unique,
    tema INTEGER,
    link TEXT default '',

    FOREIGN KEY (tema) REFERENCES tema(id) ON DELETE SET NULL,
    FOREIGN KEY (Foto) REFERENCES Foto(id) ON DELETE SET NULL
);