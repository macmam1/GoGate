using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;

namespace GoGate.WindowsClient;

public record SharedProfileDocument(string Version, List<ProfileItem> Profiles);

public static class SharedProfileConfig
{
    public static IReadOnlyList<ProfileItem> Parse(string json)
    {
        var doc = JsonSerializer.Deserialize<SharedProfileDocument>(json, new JsonSerializerOptions
        {
            PropertyNameCaseInsensitive = true
        });
        if (doc is null || !string.Equals(doc.Version, "v1", StringComparison.OrdinalIgnoreCase))
            return Array.Empty<ProfileItem>();

        return doc.Profiles ?? new List<ProfileItem>();
    }

    public static string Serialize(IEnumerable<ProfileItem> profiles)
    {
        var doc = new SharedProfileDocument("v1", profiles?.ToList() ?? new List<ProfileItem>());
        return JsonSerializer.Serialize(doc, new JsonSerializerOptions { WriteIndented = true });
    }
}
