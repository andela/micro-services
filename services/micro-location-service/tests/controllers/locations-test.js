const should = require('chai').should();
const sinon = require('sinon');
const Q = require('q');
const grpc = require('grpc');
const controller = require('../../controllers/locations_controller');
const models = require('../../models');
const producer = require('../../kafka_producer');

const locations = [
  {
    id: '-KPE-AopjdUJrbOELUuk',
    name: 'Lagos',
    time_zone: '+2',
    created_at: new Date(),
    updated_at: new Date(),
  },
  {
    id: '-KPE-IC7p2-CpGKnmUni',
    name: 'Kenya',
    time_zone: '+4',
    created_at: new Date(),
    updated_at: new Date(),
  },
];


describe('Locations controllers', () => {
  afterEach(() => {
    sinon.restore(models.Location, 'findAll');
    sinon.restore(models.Location, 'findById');
    sinon.restore(producer, 'emit');
    sinon.restore(grpc, 'load');
  });

  describe('List all Locations: Success', () => {
    beforeEach(() => {
      sinon.stub(models.Location, 'findAll', () => {
        const deferred = Q.defer();
        deferred.resolve(locations);
        return deferred.promise;
      });
    });

    it('should not return error', (done) => {
      controller.index({}, (err) => {
        should.not.exist(err);
        done();
      });
    });

    it('should return locations data', (done) => {
      controller.index({}, (err, loc) => {
        should.exist(loc);
        loc[0].name.should.equal('Lagos');
        done();
      });
    });

    it('should contain correct data count', (done) => {
      controller.index({}, (err, result) => {
        result.length.should.equal(2);
        done();
      });
    });
  });

  describe('List all Locations: failure', () => {
    beforeEach(() => {
      sinon.stub(models.Location, 'findAll', () => {
        const deferred = Q.defer();
        deferred.reject(new Error('Something happened'));
        return deferred.promise;
      });
    });


    it('should contain correct error message', (done) => {
      controller.index({}, (err, result) => {
        should.not.exist(result);
        err.message.should.equal('Something happened');
        done();
      });
    });
  });

  describe('Get a Location: success', () => {
    const call = { metadata: {}, request: { id: '-KPE-AopjdUJrbOELUuk' } };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.resolve(locations[0]);
        return deferred.promise;
      });
    });

    it('should not return error', (done) => {
      controller.show(call, (err) => {
        should.not.exist(err);
        done();
      });
    });

    it('should return data', (done) => {
      controller.show(call, (err, location) => {
        should.exist(location);
        location.name.should.equal('Lagos');
        done();
      });
    });

    it('should contain correct data', (done) => {
      controller.show(call, (err, result) => {
        result.id.should.equal('-KPE-AopjdUJrbOELUuk');
        done();
      });
    });
  });

  describe('Get a user: failure', () => {
    const call = { metadata: {}, request: { id: 1 } };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.reject(new Error('there was a problem'));
        return deferred.promise;
      });
    });

    it('should return error', (done) => {
      controller.show(call, (err) => {
        should.exist(err);
        err.message.should.equal('there was a problem');
        done();
      });
    });

    it('should not return data', (done) => {
      controller.show(call, (err, result) => {
        should.not.exist(result);
        done();
      });
    });
  });

  describe('Location not found', () => {
    const call = { metadata: {}, request: { id: '-KPE-AopjdUJrbOELUuk' } };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.resolve(null);
        return deferred.promise;
      });
    });

    it('should return error', (done) => {
      controller.show(call, (err) => {
        should.exist(err);
        err.message.should.equal('location not found');
        done();
      });
    });

    it('should not return data', (done) => {
      controller.show(call, (err, result) => {
        should.not.exist(result);
        done();
      });
    });

    it('should correct error message', (done) => {
      controller.show(call, (err) => {
        err.message.should.equal('location not found');
        done();
      });
    });
  });

  describe('Update success', () => {
    const call = { metadata: {}, request: {} };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.resolve(locations[0]);
        return deferred.promise;
      });

      sinon.stub(producer, 'emitModel', (model, data, cb) => {
        cb(null, {});
      });
    });

    it('should not return error', (done) => {
      controller.update(call, (err) => {
        should.not.exist(err);
        done();
      });
    });

    it('should  return data', (done) => {
      controller.show(call, (err, result) => {
        should.exist(result);
        result.name.should.equal('Lagos');
        done();
      });
    });
  });

  describe('Update failure', () => {
    const call = { metadata: {}, request: {} };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.reject(new Error('Error Occured'));
        return deferred.promise;
      });

      sinon.stub(producer, 'emitModel', (model, data, cb) => {
        cb(null, {});
      });
    });

    it('should return error', (done) => {
      controller.update(call, (err, result) => {
        should.exist(err);
        should.not.exist(result);
        err.message.should.equal('Error Occured');
        done();
      });
    });
  });

  describe('Create', () => {
    const call = { metadata: {}, request: { name: 'Nigeria' } };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.reject(new Error('Error Occured'));
        return deferred.promise;
      });

      sinon.stub(producer, 'emitModel', (model, data, cb) => {
        cb(null, {});
      });
    });

    it('should return data', (done) => {
      controller.create(call, (err, result) => {
        should.not.exist(err);
        should.exist(result);
        result.should.deep.equal({});
        done();
      });
    });
  });

  describe('Destroy: Success', () => {
    const call = { metadata: {}, request: {} };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.resolve(locations[0]);
        return deferred.promise;
      });

      sinon.stub(producer, 'emit', (data, cb) => {
        cb(null, {});
      });
    });

    it('should return data', (done) => {
      controller.destroy(call, (err, result) => {
        should.not.exist(err);
        should.exist(result);
        result.should.deep.equal({});
        done();
      });
    });
  });

  describe('Destroy: Failure', () => {
    const call = { metadata: {}, request: {} };
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.reject(new Error('Error occured'));
        return deferred.promise;
      });

      sinon.stub(producer, 'emit', (data, cb) => {
        cb(null, {});
      });
    });

    it('should return error', (done) => {
      controller.destroy(call, (err, result) => {
        should.exist(err);
        should.not.exist(result);
        err.message.should.equal('Error occured');
        done();
      });
    });
  });

  describe('All details', () => {
    beforeEach(() => {
      sinon.stub(models.Location, 'findAll', () => {
        const deferred = Q.defer();
        deferred.resolve(locations);
        return deferred.promise;
      });
    });

    it('should return the correct data', (done) => {
      controller.allLocationsDetails({}, (err, allLocations) => {
        allLocations.length.should.equal(2);
        allLocations[0].title.should.equal('Lagos');
        allLocations[0].count.should.equal(35);
        done();
      });
    });
  });

  describe('getLocationDetails', () => {
    beforeEach(() => {
      sinon.stub(models.Location, 'findById', () => {
        const deferred = Q.defer();
        deferred.resolve(locations[0]);
        return deferred.promise;
      });
    });

    it('should return correct count of data', (done) => {
      controller.getLocationDetails({ request: { id: '-KPE-AopjdUJrbOELUuk' } }, (err, data) => {
        data.values.length.should.equal(2);
        data.values[0].title.should.equal('D0B-SIMULATIONS');
        data.values[0].count.should.equal(13);
        done();
      });
    });
  });
});
