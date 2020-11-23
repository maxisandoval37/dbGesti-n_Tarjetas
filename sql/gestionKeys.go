package sql

import (
	"log"
)

func GenerarPKs() {
	_, err = db.Exec(`ALTER TABLE cliente ADD CONSTRAINT cliente_pk PRIMARY KEY (nrocliente);
					  ALTER TABLE tarjeta ADD CONSTRAINT tarjeta_pk PRIMARY KEY (nrotarjeta);
					  ALTER TABLE comercio ADD CONSTRAINT comercio_pk PRIMARY KEY (nrocomercio);
	                  ALTER TABLE compra ADD CONSTRAINT compra_pk PRIMARY KEY (nrooperacion);
	                  ALTER TABLE rechazo ADD CONSTRAINT rechazo_pk PRIMARY KEY (nrorechazo);
	                  ALTER TABLE cierre ADD CONSTRAINT cierre_pk PRIMARY KEY (año, mes, terminacion);
	                  ALTER TABLE cabecera ADD CONSTRAINT cabecera_pk PRIMARY KEY (nroresumen);
	                  ALTER TABLE detalle ADD CONSTRAINT detalle_pk PRIMARY KEY (nroresumen, nrolinea);
					  ALTER TABLE alerta ADD CONSTRAINT alerta_pk PRIMARY KEY (nroalerta);`)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerarFKs() {
	_, err = db.Exec(`ALTER TABLE tarjeta ADD CONSTRAINT tarjeta_nrocliente_fk FOREIGN KEY (nrocliente) REFERENCES cliente(nrocliente);
					  ALTER TABLE compra ADD CONSTRAINT compra_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE compra ADD CONSTRAINT compra_nrocomercio_fk FOREIGN KEY (nrocomercio) REFERENCES comercio(nrocomercio);
					  ALTER TABLE rechazo ADD CONSTRAINT rechazo_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE rechazo ADD CONSTRAINT rechazo_nrocomercio_fk FOREIGN KEY (nrocomercio) REFERENCES comercio(nrocomercio);
					  ALTER TABLE cabecera ADD CONSTRAINT cabecera_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE detalle ADD CONSTRAINT detalle_cabecera_fk FOREIGN KEY (nroresumen) REFERENCES cabecera(nroresumen);
					  ALTER TABLE alerta ADD CONSTRAINT alerta_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					`)			
	if err != nil {
		log.Fatal(err)
	}
}
