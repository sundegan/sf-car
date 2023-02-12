import { rental } from "./proto_gen/rental/rental_pb";
import { sfcar } from "./request";

 export namespace TripService {
    export function CreateTrip(req: rental.v1.ICreateTripRequest): Promise<rental.v1.ITripEntity> {
        return sfcar.sendRequestWithAuthRetry({
            method: 'POST',
            path: '/v1/trip',
            data: req,
            respMarshaller: rental.v1.TripEntity.fromObject,
        })
    }

    export function GetTrip(id: string): Promise<rental.v1.ITrip> {
        return sfcar.sendRequestWithAuthRetry({
            method: 'GET',
            path: `/v1/trip/${encodeURIComponent(id)}`,
            respMarshaller: rental.v1.Trip.fromObject,
        })
    }

    // 如果给定了状态参数,则返回指定状态的行程列表
    // 如果没有给定状态参数,则返回所有行程
    export function getTrips(s?: rental.v1.TripStatus): Promise<rental.v1.IGetTripsResponse> {
        let path = '/v1/trips'
        if (s) {
            path += `?status=${s}`
        }
        return sfcar.sendRequestWithAuthRetry({
            method: 'GET',
            path: path,
            respMarshaller: rental.v1.GetTripsResponse.fromObject,
        })
    }
 }