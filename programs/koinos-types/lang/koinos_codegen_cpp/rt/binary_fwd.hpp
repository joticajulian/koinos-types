#pragma once

#include <koinos/pack/rt/basetypes.hpp>
#include <koinos/pack/rt/varint.hpp>

#define KOINOS_DECLARE_PRIMITIVE_SERIALIZER( type )               \
template< typename Stream >                                       \
inline void to_binary( Stream& s, type t );                       \
template< typename Stream >                                       \
inline void from_binary( Stream& s, type& t, uint32_t depth = 0 );\

#define KOINOS_DECLARE_BASE_SERIALIZER( type )                    \
template< typename Stream >                                       \
inline void to_binary( Stream& s, const type& t );                \
template< typename Stream >                                       \
inline void from_binary( Stream& s, type& t, uint32_t depth = 0 );\

namespace koinos
{
   // Forward declaration
   template< typename T > class opaque;
} // koinos

namespace koinos::pack {

KOINOS_DECLARE_PRIMITIVE_SERIALIZER( int8_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( uint8_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( int16_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( uint16_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( int32_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( uint32_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( int64_t )
KOINOS_DECLARE_PRIMITIVE_SERIALIZER( uint64_t )

KOINOS_DECLARE_BASE_SERIALIZER( int128_t )
KOINOS_DECLARE_BASE_SERIALIZER( uint128_t )
KOINOS_DECLARE_BASE_SERIALIZER( int160_t )
KOINOS_DECLARE_BASE_SERIALIZER( uint160_t )
KOINOS_DECLARE_BASE_SERIALIZER( int256_t )
KOINOS_DECLARE_BASE_SERIALIZER( uint256_t )

KOINOS_DECLARE_PRIMITIVE_SERIALIZER( bool )

KOINOS_DECLARE_BASE_SERIALIZER( unsigned_int )
KOINOS_DECLARE_BASE_SERIALIZER( signed_int )

KOINOS_DECLARE_BASE_SERIALIZER( multihash )
KOINOS_DECLARE_BASE_SERIALIZER( multihash_vector )

KOINOS_DECLARE_BASE_SERIALIZER( block_height_type )
KOINOS_DECLARE_BASE_SERIALIZER( timestamp_type )

KOINOS_DECLARE_BASE_SERIALIZER( std::string )
KOINOS_DECLARE_BASE_SERIALIZER( variable_blob )

template< typename Stream, size_t N >
inline void to_binary( Stream& s, const fixed_blob< N >& v );
template< typename Stream, size_t N >
inline void from_binary( Stream& s, fixed_blob< N >& v, uint32_t depth = 0 );

template< typename Stream, typename T >
inline void to_binary( Stream& s, const std::vector< T >& v );
template< typename Stream, typename T >
inline void from_binary( Stream& s, std::vector< T >& v, uint32_t depth = 0 );

template< typename Stream, typename T, size_t N >
inline void to_binary( Stream& s, const std::array< T, N >& v );
template< typename Stream, typename T, size_t N >
inline void from_binary( Stream& s, std::array< T, N >& v, uint32_t depth = 0 );

template< typename Stream, typename... T >
inline void to_binary( Stream& s, const std::variant< T... >& v );
template< typename Stream, typename... T >
inline void from_binary( Stream& s, std::variant< T... >& v, uint32_t depth = 0 );

template< typename Stream, typename T >
inline void to_binary( Stream& s, const std::optional< T >& v );
template< typename Stream, typename T >
inline void from_binary( Stream& s, std::optional< T >& v, uint32_t depth = 0 );

template< typename Stream, typename T >
inline void to_binary( Stream& s, const opaque< T >& v );
template< typename Stream, typename T >
inline void from_binary( Stream& s, opaque< T >& v, uint32_t depth = 0 );

template< typename Stream, typename T >
inline void to_binary( Stream& s, const T& v );
template< typename Stream, typename T >
inline void from_binary( Stream& s, T& v, uint32_t depth = 0 );

template< typename T >
inline void to_variable_blob( variable_blob& v, const T& t, bool append = false );
template< typename T >
inline variable_blob to_variable_blob( const T& t );

template< typename T >
inline void from_variable_blob( const variable_blob& v, T& t );
template< typename T >
inline T from_variable_blob( const variable_blob& v );

template< typename T, size_t N >
inline void from_fixed_blob( const fixed_blob< N >& f, T& t );
template< typename T, size_t N >
inline T from_fixed_blob( const fixed_blob< N >& f );

template< typename T >
inline void to_c_str( char* c, size_t l, const T& t );
template< typename T >
inline void from_c_str( const char* c, size_t l, T& t );
template< typename T >
inline T from_c_str( const char* c, size_t l );

} // koinos::pack
