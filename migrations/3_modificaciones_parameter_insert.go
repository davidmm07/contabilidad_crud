package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		AdicionApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "adicion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		ReduccionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "reduccion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		SuspencionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "suspencion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AdicionApropiacion,
			ReduccionApropiacion,
			SuspencionApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		AdicionApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "adicion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
		}
		ReduccionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "reduccion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
		}
		SuspencionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "suspencion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
		}

		parameters := []interface{}{
			AdicionApropiacion,
			ReduccionApropiacion,
			SuspencionApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}